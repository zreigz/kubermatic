package graph

import (
	"context"
	"testing"
	"time"
)

type testTask struct {
	done  bool
	sleep time.Duration
}

func (t *testTask) Execute() error {
	go func() {
		time.Sleep(t.sleep)
		t.done = true
	}()

	return nil
}

func (t *testTask) Done() (bool, error) {
	return t.done, nil
}

func TestResolver_AddTask(t *testing.T) {
	r := NewResolver()

	tasks := []struct {
		name         string
		dependencies []string
	}{
		{
			name: "root_ca",
		},
		{
			name:         "apiserver_tls",
			dependencies: []string{"root_ca"},
		},
		{
			name:         "apiserver_kubelet_client",
			dependencies: []string{"root_ca"},
		},
		{
			name:         "apiserver_deployment",
			dependencies: []string{"root_ca", "apiserver_kubelet_client", "apiserver_tls"},
		},
	}

	for _, task := range tasks {
		if err := r.AddTask(task.name, &testTask{}, task.dependencies...); err != nil {
			t.Fatalf("failed to add task '%s': %v", task.name, err)
		}
	}

	if err := PrintGraph(r.dag, "task-graph.svg"); err != nil {
		t.Fatal(err)
	}
}

func TestResolver_Resolve(t *testing.T) {
	r := NewResolver()

	tasks := []struct {
		name         string
		dependencies []string
		sleep        time.Duration
	}{
		{
			name:  "root_ca",
			sleep: 100 * time.Millisecond,
		},
		{
			name:         "apiserver_tls",
			dependencies: []string{"root_ca"},
		},
		{
			name:         "apiserver_kubelet_client",
			dependencies: []string{"root_ca"},
		},
		{
			name:         "apiserver_deployment",
			dependencies: []string{"root_ca", "apiserver_kubelet_client", "apiserver_tls"},
		},
	}

	for _, task := range tasks {
		if err := r.AddTask(task.name, &testTask{}, task.dependencies...); err != nil {
			t.Fatalf("failed to add task '%s': %v", task.name, err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	if err := r.Resolve(ctx); err != nil {
		t.Fatalf("failed to resolve graph: %v", err)
	}
}
