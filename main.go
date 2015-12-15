package main

import (
	"github.com/paullius/go-graphs-/directed"
	"github.com/paullius/go-graphs-/undirected"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree"
	"github.com/paullius/go-graphs-/weighted/shortestPaths"
)

func main() {

	undirected.DepthFirstExample(0)
	undirected.ConnectedComponentsExample()
	undirected.DepthFirstPathsExample(0)
	undirected.BreadthFirstPathsExample(0)

	var dg = directed.NewDigraph(5)
	dg.AddEdge(1, 2)
	dg.AddEdge(1, 3)
	dg.AddEdge(2, 3)

	minimumSpanningTree.PrimMstExample()

	shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})

}

//go install github.com/paullius/go-graphs-
//go test github.com/paullius/go-graphs-\undirected
