package directed

import (
	"fmt"
)

func DirectedDepthFirstExample(from int) {

	g := createDirectedDepthFirstGraph()
	dfp := g.DirecteDepthFirst(from)

	dfp.Print(from)
	fmt.Println()
}

func createDirectedDepthFirstGraph() Digraph {
	g := NewDigraph(13)
	g.AddEdge(4, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 2)
	g.AddEdge(6, 0)
	g.AddEdge(0, 1)
	g.AddEdge(2, 0)
	g.AddEdge(11, 12)
	g.AddEdge(12, 9)
	g.AddEdge(9, 10)
	g.AddEdge(9, 11)
	g.AddEdge(8, 9)
	g.AddEdge(10, 12)
	g.AddEdge(11, 4)
	g.AddEdge(4, 3)
	g.AddEdge(3, 5)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(5, 4)
	g.AddEdge(0, 5)
	g.AddEdge(6, 4)
	g.AddEdge(6, 9)
	g.AddEdge(7, 6)

	return g
}
