package shortestPaths

import (
	"github.com/paullius/go-graphs-/weighted"
)

type EdgeWeightedDigraph struct {
	v, e int
	adj  [][]weighted.Edge
}

func NewEdgeWeightedDigraph(v int) EdgeWeightedDigraph {
	g := EdgeWeightedDigraph{
		v,
		0,
		make([][]weighted.Edge, v),
	}
	return g
}

func (g *EdgeWeightedDigraph) AddEdge(e weighted.Edge) {
	g.adj[e.From()] = append(g.adj[e.From()], e)

	g.e++
}

func (g *EdgeWeightedDigraph) AdjacentTo(v int) []weighted.Edge {
	return g.adj[v]
}

func (g *EdgeWeightedDigraph) Edges() []weighted.Edge {
	var b []weighted.Edge

	for v := 0; v < g.v; v++ {
		for _, e := range g.adj[v] {
			b = append(b, e)
		}
	}

	return b
}

func (g *EdgeWeightedDigraph) V() int {
	return g.v
}
