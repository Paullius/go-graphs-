package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/nonWeighted"
	"github.com/paullius/go-graphs-/nonWeighted/undirected"
)

func ConnectedComponentsExample() {

	g := createConnectedComponentsGraph()
	cc := undirected.NewConnectedComponents(g)

	cc.Print()
	fmt.Println()
}

func createConnectedComponentsGraph() nonWeighted.NonWeightedGraph {
	g := nonWeighted.NewNonWeightedGraph(13)
	g.AddUndirectedEdge(0, 6)
	g.AddUndirectedEdge(0, 2)
	g.AddUndirectedEdge(0, 1)
	g.AddUndirectedEdge(0, 5)
	g.AddUndirectedEdge(3, 5)
	g.AddUndirectedEdge(3, 4)
	g.AddUndirectedEdge(4, 5)
	g.AddUndirectedEdge(4, 6)
	g.AddUndirectedEdge(7, 8)
	g.AddUndirectedEdge(9, 10)
	g.AddUndirectedEdge(9, 11)
	g.AddUndirectedEdge(9, 12)
	g.AddUndirectedEdge(11, 12)
	return g
}
