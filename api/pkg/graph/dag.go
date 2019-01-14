package graph

import (
	"fmt"
	"os/exec"
	"strings"
)

type Node struct {
	value interface{}
	name  string

	children []*Node
}

type DAG struct {
	nodes map[string]*Node
}

func NewDAG() *DAG {
	d := &DAG{}
	d.nodes = make(map[string]*Node)
	return d
}

func (d *DAG) AddVertex(name string, v interface{}) *Node {
	node := &Node{name: name, value: v}
	d.nodes[name] = node
	return node
}

func (d *DAG) AddEdge(from, to string) error {
	fromNode, exists := d.nodes[from]
	if !exists {
		return fmt.Errorf("from node '%s' does not exist", from)
	}

	toNode, exists := d.nodes[to]
	if !exists {
		return fmt.Errorf("to node '%s' does not exist", to)
	}

	fromNode.children = append(fromNode.children, toNode)

	return nil
}

// GetDotGraph returns the graph as dot code
func GetDotGraph(d *DAG) string {
	s := &strings.Builder{}
	s.WriteString("digraph D {\n")
	for _, node := range d.nodes {
		for _, childNode := range node.children {
			s.WriteString(fmt.Sprintf("  %s -> %s\n", node.name, childNode.name))
		}
	}
	s.WriteString("}\n")
	return s.String()
}

func PrintGraph(d *DAG, filename string) error {
	s := GetDotGraph(d)

	cmd := exec.Command("/usr/bin/dot", "-o", filename, "-T", "svg")
	// Pipe in the dot text
	in := strings.NewReader(s)
	cmd.Stdin = in

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to call dot: %v. out: %s", err, string(out))
	}

	return nil
}
