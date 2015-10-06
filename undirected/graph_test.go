package undirected_test

import (
	"github.com/paullius/go-graphs-/undirected"
	"testing"
)

func TestNewGraphVertices(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{2, 2},
		{0, 0},
	}
	for _, c := range cases {
		graph := undirected.NewGraph(c.in)
		got := graph.Vertices()
		if got != c.want {
			t.Errorf("NewGraph(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestNewGraphEdges(t *testing.T) {
	g := createTestGraph()
	if g.Edges() != 7 {
		t.Errorf("Edges = %v, want %v", g.Edges(), 7)
	}
}

func createTestGraph() undirected.Graph {
	g := undirected.NewGraph(7)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 6)
	g.AddEdge(4, 6)
	return g
}
