package shortestPaths

import (
	"fmt"
	"github.com/paullius/go-graphs-/weighted"
)

func DijkstrasExample(from int, to []int) {

	g := createShortestPathGraph()
	sp := NewDijkstraShortestPaths(g, from)

	sp.PrintShortestPath(to)
	fmt.Println()
}

func createShortestPathGraph() EdgeWeightedDigraph {
	g := NewEdgeWeightedDigraph(8)
	g.AddEdge(weighted.NewEdge(4, 5, 0.35))
	g.AddEdge(weighted.NewEdge(5, 4, 0.35))
	g.AddEdge(weighted.NewEdge(4, 7, 0.37))
	g.AddEdge(weighted.NewEdge(5, 7, 0.28))
	g.AddEdge(weighted.NewEdge(7, 5, 0.28))
	g.AddEdge(weighted.NewEdge(5, 1, 0.32))
	g.AddEdge(weighted.NewEdge(0, 4, 0.38))
	g.AddEdge(weighted.NewEdge(0, 2, 0.26))
	g.AddEdge(weighted.NewEdge(7, 3, 0.39))
	g.AddEdge(weighted.NewEdge(1, 3, 0.29))
	g.AddEdge(weighted.NewEdge(2, 7, 0.34))
	g.AddEdge(weighted.NewEdge(6, 2, 0.40))
	g.AddEdge(weighted.NewEdge(3, 6, 0.52))
	g.AddEdge(weighted.NewEdge(6, 0, 0.58))
	g.AddEdge(weighted.NewEdge(6, 4, 0.93))

	return g
}
