package minimumSpanningTree

import (
	"fmt"
	"github.com/paullius/go-graphs-/weighted"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree"
)

func PrimMstExample() {

	g := createEdgeWeightedGraph()
	sp := minimumSpanningTree.NewPrimMst(g)
	sp.PrintMinimumSpanningTree()
	fmt.Println()
}

func createEdgeWeightedGraph() weighted.EdgeWeightedGraph {
	g := weighted.NewEdgeWeightedGraph(8)
	g.AddEdge(weighted.NewEdge(4, 5, 0.35))
	g.AddEdge(weighted.NewEdge(4, 7, 0.37))
	g.AddEdge(weighted.NewEdge(5, 7, 0.28))
	g.AddEdge(weighted.NewEdge(0, 7, 0.16))
	g.AddEdge(weighted.NewEdge(1, 5, 0.32))
	g.AddEdge(weighted.NewEdge(0, 4, 0.38))
	g.AddEdge(weighted.NewEdge(2, 3, 0.17))
	g.AddEdge(weighted.NewEdge(1, 7, 0.19))
	g.AddEdge(weighted.NewEdge(0, 2, 0.26))
	g.AddEdge(weighted.NewEdge(1, 2, 0.36))
	g.AddEdge(weighted.NewEdge(1, 3, 0.29))
	g.AddEdge(weighted.NewEdge(2, 7, 0.34))
	g.AddEdge(weighted.NewEdge(6, 2, 0.40))
	g.AddEdge(weighted.NewEdge(3, 6, 0.52))
	g.AddEdge(weighted.NewEdge(6, 0, 0.58))
	g.AddEdge(weighted.NewEdge(6, 4, 0.93))

	return g
}
