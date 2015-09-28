package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"
)

func main() {
	var g = undirected.CreateGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	var dfs = g.DepthFirstSearch(2)
	var dfp = g.DepthFirstPathsSearch(2)
	var bfp = g.BreadthFirstPathsSearch(1)

	fmt.Println(g.V)
	fmt.Println(dfs.Count)
	fmt.Println(dfp.S)
	fmt.Println(bfp.S)
	fmt.Println(bfp.HasPathTo(2))
	fmt.Println(bfp.HasPathTo(4))
}

//go install github.com/paullius/go-graphs-
