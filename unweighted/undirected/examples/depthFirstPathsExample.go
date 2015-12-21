package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/unweighted"
	"github.com/paullius/go-graphs-/unweighted/undirected"
)

func DepthFirstPathsExample(from int) {

	g := createDepthFirstPathsGraph()
	dfp := undirected.NewDepthFirstPaths(g, from)

	dfp.Print(from)
	fmt.Println()
}

func createDepthFirstPathsGraph() unweighted.NonWeightedGraph {
	g := unweighted.NewNonWeightedGraph(13)
	g.AddUndirectedEdge(0, 6)
	g.AddUndirectedEdge(0, 2)
	g.AddUndirectedEdge(0, 1)
	g.AddUndirectedEdge(0, 5)
	g.AddUndirectedEdge(3, 5)
	g.AddUndirectedEdge(3, 4)
	g.AddUndirectedEdge(4, 5)
	g.AddUndirectedEdge(4, 6)
	g.AddUndirectedEdge(7, 8)
	g.AddUndirectedEdge(9, 10)
	g.AddUndirectedEdge(9, 11)
	g.AddUndirectedEdge(9, 12)
	g.AddUndirectedEdge(11, 12)
	return g
}
