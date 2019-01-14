package graph

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

type Task interface {
	Execute() error
	Done() (bool, error)
}

type Resolver struct {
	lock *sync.Mutex
	done map[string]bool
	dag  *DAG
}

func NewResolver() *Resolver {
	return &Resolver{
		dag:  NewDAG(),
		done: map[string]bool{},
		lock: &sync.Mutex{},
	}
}

func (r *Resolver) AddTask(name string, t Task, dependencies ...string) error {
	r.dag.AddVertex(name, t)

	for _, dep := range dependencies {
		if err := r.dag.AddEdge(name, dep); err != nil {
			return fmt.Errorf("failed to add edge from='%s' to='%s': %v", name, dep, err)
		}
	}

	return nil
}

func (r *Resolver) Resolve(ctx context.Context) error {
	var reachedTimeout bool
	go func() {
		<-ctx.Done()
		reachedTimeout = true
	}()

	for _, node := range r.dag.nodes {
		if reachedTimeout {
			return fmt.Errorf("reached timeout while solving the tasks")
		}

		if err := r.solveNode(node, ctx); err != nil {
			return fmt.Errorf("failed to solve node '%s': %v", node.name, err)
		}
	}
	return nil
}

func (r *Resolver) canBeExecuted(currentNode *Node) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, dep := range currentNode.children {
		if !r.done[dep.name] {
			return false
		}
	}

	return true
}

func (r *Resolver) solved(currentNode *Node) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	return r.done[currentNode.name]
}

func (r *Resolver) solveNode(currentNode *Node, ctx context.Context) error {
	lock := &sync.Mutex{}
	errorBuilder := &strings.Builder{}
	wg := &sync.WaitGroup{}

	if r.solved(currentNode) {
		return nil
	}

	for _, child := range currentNode.children {

		wg.Add(1)
		go func(node *Node) {
			defer wg.Done()
			if err := r.solveNode(node, ctx); err != nil {
				lock.Lock()
				defer lock.Unlock()
				errorBuilder.WriteString(err.Error() + "\n")
			}
		}(child)
		wg.Wait()

	}

	if r.canBeExecuted(currentNode) {
		task := currentNode.value.(Task)
		if err := task.Execute(); err != nil {
			return fmt.Errorf("failed to execute task: %v", err)
		}

		if err := wait.PollImmediateUntil(25*time.Millisecond, task.Done, ctx.Done()); err != nil {
			// We won't populate an error in case we hit the timeout, as the top level knows about it
			if err == wait.ErrWaitTimeout {
				return nil
			}
			return fmt.Errorf("failed to check if the task is done: %v", err)
		}

		// Use extra func to minimize the locking of the done map
		func() {
			r.lock.Lock()
			defer r.lock.Unlock()
			r.done[currentNode.name] = true
		}()
		log.Printf("Solved '%s'", currentNode.name)
	}

	if errorBuilder.Len() != 0 {
		return fmt.Errorf(errorBuilder.String())
	}
	return nil
}
