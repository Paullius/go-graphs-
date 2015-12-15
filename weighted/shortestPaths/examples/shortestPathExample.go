package shortestPaths

import (
	"fmt"
	"github.com/paullius/go-graphs-/weighted"
	"github.com/paullius/go-graphs-/weighted/shortestPaths"
)

func DijkstrasExample(from int, to []int) {

	g := createShortestPathGraph()
	sp := shortestPaths.NewDijkstraShortestPaths(g, from)

	sp.PrintShortestPath(to)
	fmt.Println()
}

func createShortestPathGraph() weighted.EdgeWeightedGraph {
	g := weighted.NewEdgeWeightedGraph(8)
	g.AddDirectedEdge(weighted.NewEdge(4, 5, 0.35))
	g.AddDirectedEdge(weighted.NewEdge(5, 4, 0.35))
	g.AddDirectedEdge(weighted.NewEdge(4, 7, 0.37))
	g.AddDirectedEdge(weighted.NewEdge(5, 7, 0.28))
	g.AddDirectedEdge(weighted.NewEdge(7, 5, 0.28))
	g.AddDirectedEdge(weighted.NewEdge(5, 1, 0.32))
	g.AddDirectedEdge(weighted.NewEdge(0, 4, 0.38))
	g.AddDirectedEdge(weighted.NewEdge(0, 2, 0.26))
	g.AddDirectedEdge(weighted.NewEdge(7, 3, 0.39))
	g.AddDirectedEdge(weighted.NewEdge(1, 3, 0.29))
	g.AddDirectedEdge(weighted.NewEdge(2, 7, 0.34))
	g.AddDirectedEdge(weighted.NewEdge(6, 2, 0.40))
	g.AddDirectedEdge(weighted.NewEdge(3, 6, 0.52))
	g.AddDirectedEdge(weighted.NewEdge(6, 0, 0.58))
	g.AddDirectedEdge(weighted.NewEdge(6, 4, 0.93))

	return g
}
