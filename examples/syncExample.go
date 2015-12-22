package examples

import (
	"github.com/paullius/go-graphs-/unweighted/directed/examples"
	"github.com/paullius/go-graphs-/unweighted/undirected/examples"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree/examples"
	"github.com/paullius/go-graphs-/weighted/shortestPaths/examples"
)

func SyncRun() {
	undirected.DepthFirstExample(0)
	undirected.ConnectedComponentsExample()
	undirected.DepthFirstPathsExample(0)
	undirected.BreadthFirstPathsExample(0)

	directed.DirectedDepthFirstExample(2)

	minimumSpanningTree.PrimMstExample()

	shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})
}
