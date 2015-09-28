package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"
)

func main() {
	var g undirected.Graph = undirected.CreateGraph(5)
	undirected.AddEdge(1, 2, g)
	undirected.AddEdge(1, 3, g)
	undirected.AddEdge(2, 3, g)
	var dfs undirected.DepthFirst = undirected.DepthFirstSearch(g, 2)
	var dfps undirected.DepthFirstPaths = undirected.DepthFirstPathsSearch(g, 2)
	var bfps undirected.BreadthFirstPaths = undirected.BreadthFirstPathsSearch(g, 1)

	fmt.Println(g.V)
	fmt.Println(dfs.Count)
	fmt.Println(dfps.S)
	fmt.Println(bfps.S)

}

//go install github.com/paullius/go-graphs-
