package main

import (
	"github.com/paullius/go-graphs-/nonWeighted/directed/examples"
	"github.com/paullius/go-graphs-/nonWeighted/undirected/examples"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree/examples"
	"github.com/paullius/go-graphs-/weighted/shortestPaths/examples"
)

func main() {

	undirected.DepthFirstExample(0)
	undirected.ConnectedComponentsExample()
	undirected.DepthFirstPathsExample(0)
	undirected.BreadthFirstPathsExample(0)

	directed.DirectedDepthFirstExample(2)

	minimumSpanningTree.PrimMstExample()

	shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})

}

//go run %GOPATH%\src\github.com\Paullius\go-graphs-\main.go
//go install github.com/paullius/go-graphs-
//go test github.com/paullius/go-graphs-\nonWeighted
