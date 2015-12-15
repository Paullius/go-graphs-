package undirected

import (
	"fmt"
)

func DepthFirstExample(from int) {

	g := createDepthFirstGraph()
	df := g.DepthFirst(from)

	df.Print(from)
	fmt.Println()
}

func createDepthFirstGraph() Graph {
	g := NewGraph(13)
	g.AddEdge(0, 5)
	g.AddEdge(4, 3)
	g.AddEdge(0, 1)
	g.AddEdge(9, 12)
	g.AddEdge(6, 4)
	g.AddEdge(5, 4)
	g.AddEdge(0, 2)
	g.AddEdge(11, 12)
	g.AddEdge(9, 10)
	g.AddEdge(0, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(5, 3)

	return g
}
