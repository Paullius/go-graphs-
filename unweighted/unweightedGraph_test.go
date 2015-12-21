package unweighted_test

import (
	"github.com/paullius/go-graphs-/unweighted"
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
		graph := unweighted.NewUnweightedGraph(c.in)
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

func createTestGraph() unweighted.UnweightedGraph {
	g := unweighted.NewUnweightedGraph(7)
	g.AddUndirectedEdge(0, 2)
	g.AddUndirectedEdge(0, 3)
	g.AddUndirectedEdge(1, 3)
	g.AddUndirectedEdge(1, 4)
	g.AddUndirectedEdge(1, 5)
	g.AddUndirectedEdge(2, 6)
	g.AddUndirectedEdge(4, 6)
	return g
}
