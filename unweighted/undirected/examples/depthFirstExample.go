package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/unweighted"
	"github.com/paullius/go-graphs-/unweighted/undirected"
)

func DepthFirstExample(from int) {

	g := createDepthFirstGraph()
	df := undirected.NewDepthFirst(g, from)

	df.Print(from)
	fmt.Println()
}

func createDepthFirstGraph() unweighted.UnweightedGraph {
	g := unweighted.NewUnweightedGraph(13)
	g.AddUndirectedEdge(0, 5)
	g.AddUndirectedEdge(4, 3)
	g.AddUndirectedEdge(0, 1)
	g.AddUndirectedEdge(9, 12)
	g.AddUndirectedEdge(6, 4)
	g.AddUndirectedEdge(5, 4)
	g.AddUndirectedEdge(0, 2)
	g.AddUndirectedEdge(11, 12)
	g.AddUndirectedEdge(9, 10)
	g.AddUndirectedEdge(0, 6)
	g.AddUndirectedEdge(7, 8)
	g.AddUndirectedEdge(9, 11)
	g.AddUndirectedEdge(5, 3)

	return g
}
