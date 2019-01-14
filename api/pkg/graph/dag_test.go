package graph

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDAG_AddVertex(t *testing.T) {
	d := NewDAG()
	d.AddVertex("A", 1)
	d.AddVertex("AA", 21)
	d.AddVertex("AB", 22)

	d.AddEdge("A", "AA")
	d.AddEdge("A", "AB")

	rootNode := d.nodes["A"]
	if rootNode.children[0].value.(int) != 21 {
		t.Fatalf("expected first child of node A(1) to contain node AA(21) but got %d", rootNode.children[0].value.(int))
	}
	if rootNode.children[1].value.(int) != 22 {
		t.Fatalf("expected first child of node A(1) to contain node AB(22) but got %d", rootNode.children[1].value.(int))
	}
}

func TestPrintGraph(t *testing.T) {
	d := NewDAG()
	d.AddVertex("A", 1)
	d.AddVertex("AA", 21)
	d.AddVertex("AB", 22)
	d.AddVertex("AC", 23)
	d.AddVertex("AD", 24)

	d.AddVertex("ABA", 321)
	d.AddVertex("ABB", 322)

	d.AddEdge("A", "AA")
	d.AddEdge("A", "AB")
	d.AddEdge("A", "AC")
	d.AddEdge("A", "AD")

	d.AddEdge("AB", "ABA")
	d.AddEdge("AB", "ABB")

	expected, err := ioutil.ReadFile("expected_graph.svg")
	if err != nil {
		t.Fatalf("failed to read expected_graph.svg: %v", err)
	}

	if err := PrintGraph(d, "graph.svg"); err != nil {
		t.Fatalf("failed to print graph: %v", err)
	}
	defer func() {
		os.Remove("graph.svg")
	}()

	got, err := ioutil.ReadFile("graph.svg")
	if err != nil {
		t.Fatalf("failed to read graph.svg: %v", err)
	}

	if string(expected) != string(got) {
		t.Fatalf("Generated graph differs from expected.\n Expected: \n%s\n\nGot: \n\n%s", string(expected), string(got))
	}
}
