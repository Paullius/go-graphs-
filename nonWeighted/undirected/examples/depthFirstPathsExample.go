package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/nonWeighted/undirected"
)

func DepthFirstPathsExample(from int) {

	g := createDepthFirstPathsGraph()
	dfp := g.DepthFirstPaths(from)

	dfp.Print(from)
	fmt.Println()
}

func createDepthFirstPathsGraph() undirected.Graph {
	g := undirected.NewGraph(13)
	g.AddEdge(0, 6)
	g.AddEdge(0, 2)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 10)
	g.AddEdge(9, 11)
	g.AddEdge(9, 12)
	g.AddEdge(11, 12)
	return g
}