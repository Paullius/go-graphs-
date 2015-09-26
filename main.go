 package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/undirected"	
)


 func main() {
 	var g undirected.Graph = undirected.Create(5)

 	fmt.Println(g.E)
}
