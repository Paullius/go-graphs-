package shortest

type EdgeWeightedDigraph struct {
	v, e int
	adj  [][]DirectedEdge
}

func NewEdgeWeightedGraph(v int) EdgeWeightedDigraph {
	g := EdgeWeightedDigraph{
		v,
		0,
		make([][]DirectedEdge, v),
	}
	return g
}

func (g *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	g.adj[e.from()] = append(e.from(), e)
	g.e++
}

func (g *EdgeWeightedDigraph) AdjacentTo(v int) []DirectedEdge {
	return g.adj[v]
}

func (g *EdgeWeightedDigraph) Edges() []DirectedEdge {
	var b []DirectedEdge

	for v := 0; v < g.v; v++ {
		for _, e := range g.adj[v] {
			b = append(b, e)
		}
	}

	return b
}
