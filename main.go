 package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"	
)


 func main() {
 	var g undirected.Graph = undirected.CreateGraph(5)
 	var dfs undirected.DepthFirst = undirected.DepthFirstSearch(g,2)

 	fmt.Println(g.V)
 	fmt.Println(dfs.Count)
}

//go install github.com/paullius/go-graphs-