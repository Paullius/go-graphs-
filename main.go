package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"
)

func main() {
	var g = undirected.NewGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	var dfs = g.DepthFirst(2)
	var dfp = g.DepthFirstPaths(2)
	var bfp = g.BreadthFirstPaths(1)

	fmt.Println(g.Vertices())
	fmt.Println(dfs.Count)
	fmt.Println(dfp.S)
	fmt.Println(bfp.S)
	fmt.Println(bfp.HasPathTo(2))
	fmt.Println(bfp.HasPathTo(4))
}

//go install github.com/paullius/go-graphs-
