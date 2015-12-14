package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/directed"
	"github.com/paullius/go-graphs-/undirected"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree"
	"github.com/paullius/go-graphs-/weighted/shortestPaths"
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

	minimumSpanningTree.PrimMstExample()

	shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})

}

//go install github.com/paullius/go-graphs-
//go test github.com/paullius/go-graphs-\undirected
