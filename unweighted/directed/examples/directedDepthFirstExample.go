package directed

import (
	"fmt"
	"github.com/paullius/go-graphs-/unweighted"
	"github.com/paullius/go-graphs-/unweighted/directed"
)

func DirectedDepthFirstExample(from int) {

	g := createDirectedDepthFirstGraph()
	dfp := directed.NewDirecteDepthFirst(g, from)

	dfp.Print(from)
	fmt.Println()
}

func createDirectedDepthFirstGraph() unweighted.UnweightedGraph {
	g := unweighted.NewUnweightedGraph(13)
	g.AddDirectedEdge(4, 2)
	g.AddDirectedEdge(2, 3)
	g.AddDirectedEdge(3, 2)
	g.AddDirectedEdge(6, 0)
	g.AddDirectedEdge(0, 1)
	g.AddDirectedEdge(2, 0)
	g.AddDirectedEdge(11, 12)
	g.AddDirectedEdge(12, 9)
	g.AddDirectedEdge(9, 10)
	g.AddDirectedEdge(9, 11)
	g.AddDirectedEdge(8, 9)
	g.AddDirectedEdge(10, 12)
	g.AddDirectedEdge(11, 4)
	g.AddDirectedEdge(4, 3)
	g.AddDirectedEdge(3, 5)
	g.AddDirectedEdge(7, 8)
	g.AddDirectedEdge(8, 7)
	g.AddDirectedEdge(5, 4)
	g.AddDirectedEdge(0, 5)
	g.AddDirectedEdge(6, 4)
	g.AddDirectedEdge(6, 9)
	g.AddDirectedEdge(7, 6)

	return g
}
