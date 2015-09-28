 package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"	
)


 func main() {
 	var g undirected.Graph = undirected.CreateGraph(5)
 	undirected.AddEdge(1,2,g)
 	var dfs undirected.DepthFirst = undirected.DepthFirstSearch(g,2)
 	var dfps undirected.DepthFirstPaths = undirected.DepthFirstPathsSearch(g,2)
 	var bfps undirected.BreadthFirstPaths = undirected.BreadthFirstPathsSearch(g,2)

 	fmt.Println(g.V)
 	fmt.Println(dfs.Count)
 	fmt.Println(dfps.S)
 	fmt.Println(bfps.S)
}

//go install github.com/paullius/go-graphs-