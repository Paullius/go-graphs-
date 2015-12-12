package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/directed"
	"github.com/paullius/go-graphs-/shortest"
	"github.com/paullius/go-graphs-/undirected"
	"github.com/paullius/go-graphs-/weighted"
)

func main() {
	var g = undirected.NewGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	fmt.Printf("Count of self loops: %d\n", g.NumberOfSelfLoops())

	var dfs = g.DepthFirst(2)
	var dfp = g.DepthFirstPaths(2)
	var bfp = g.BreadthFirstPaths(1)

	fmt.Printf("Count of vertices: %d\n", g.Vertices())
	fmt.Printf("Count of vertices connected to s: %d\n", dfs.ConnectedCount())
	fmt.Println(dfp.HasPathTo(1))
	fmt.Println(bfp.HasPathTo(1))
	fmt.Println(bfp.HasPathTo(2))
	fmt.Println(bfp.HasPathTo(4))

	var dg = directed.NewDigraph(5)
	dg.AddEdge(1, 2)
	dg.AddEdge(1, 3)
	dg.AddEdge(2, 3)

	var we = weighted.NewEdge(1, 2, 1.2)
	var wg = weighted.NewEdgeWeightedGraph(3)
	wg.AddEdge(we)
	var wes = wg.Edges()
	fmt.Println(wes[0].Weight())
	var lpm = weighted.NewLazyPrimMst(wg)
	fmt.Println(lpm.Edges())

	shortestPath(2, []int{3, 5})
}

func shortestPath(from int, to []int) {

	g := createShortestPathGraph()
	sp := shortest.NewDijkstraShortestPaths(g, from)
	for _, t := range to {
		sp.PrintShortestPath(t)
	}
}

func createShortestPathGraph() shortest.EdgeWeightedDigraph {
	g := shortest.NewEdgeWeightedDigraph(8)
	g.AddEdge(shortest.NewEdge(4, 5, 0.35))
	g.AddEdge(shortest.NewEdge(5, 4, 0.35))
	g.AddEdge(shortest.NewEdge(4, 7, 0.37))
	g.AddEdge(shortest.NewEdge(5, 7, 0.28))
	g.AddEdge(shortest.NewEdge(7, 5, 0.28))
	g.AddEdge(shortest.NewEdge(5, 1, 0.32))
	g.AddEdge(shortest.NewEdge(0, 4, 0.38))
	g.AddEdge(shortest.NewEdge(0, 2, 0.26))
	g.AddEdge(shortest.NewEdge(7, 3, 0.39))
	g.AddEdge(shortest.NewEdge(1, 3, 0.29))
	g.AddEdge(shortest.NewEdge(2, 7, 0.34))
	g.AddEdge(shortest.NewEdge(6, 2, 0.40))
	g.AddEdge(shortest.NewEdge(3, 6, 0.52))
	g.AddEdge(shortest.NewEdge(6, 0, 0.58))
	g.AddEdge(shortest.NewEdge(6, 4, 0.93))

	return g
}

//go install github.com/paullius/go-graphs-
//go test github.com/paullius/go-graphs-\undirected
