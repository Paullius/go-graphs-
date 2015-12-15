package minimumSpanningTree

import (
	"github.com/paullius/go-graphs-/weighted"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree"
)

func LazyPrimMstExample() {

	g := createLazyEdgeWeightedGraph()
	sp := minimumSpanningTree.NewLazyPrimMst(g)
	sp.PrintMinimumSpanningTree()
}

func createLazyEdgeWeightedGraph() weighted.EdgeWeightedGraph {
	g := weighted.NewEdgeWeightedGraph(8)
	g.AddUndirectedEdge(weighted.NewEdge(4, 5, 0.35))
	g.AddUndirectedEdge(weighted.NewEdge(4, 7, 0.37))
	g.AddUndirectedEdge(weighted.NewEdge(5, 7, 0.28))
	g.AddUndirectedEdge(weighted.NewEdge(0, 7, 0.16))
	g.AddUndirectedEdge(weighted.NewEdge(1, 5, 0.32))
	g.AddUndirectedEdge(weighted.NewEdge(0, 4, 0.38))
	g.AddUndirectedEdge(weighted.NewEdge(2, 3, 0.17))
	g.AddUndirectedEdge(weighted.NewEdge(1, 7, 0.19))
	g.AddUndirectedEdge(weighted.NewEdge(0, 2, 0.26))
	g.AddUndirectedEdge(weighted.NewEdge(1, 2, 0.36))
	g.AddUndirectedEdge(weighted.NewEdge(1, 3, 0.29))
	g.AddUndirectedEdge(weighted.NewEdge(2, 7, 0.34))
	g.AddUndirectedEdge(weighted.NewEdge(6, 2, 0.40))
	g.AddUndirectedEdge(weighted.NewEdge(3, 6, 0.52))
	g.AddUndirectedEdge(weighted.NewEdge(6, 0, 0.58))
	g.AddUndirectedEdge(weighted.NewEdge(6, 4, 0.93))

	return g
}
