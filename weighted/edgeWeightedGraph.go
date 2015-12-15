package weighted

type EdgeWeightedGraph struct {
	v, e int
	adj  [][]Edge
}

func NewEdgeWeightedGraph(v int) EdgeWeightedGraph {
	g := EdgeWeightedGraph{
		v,
		0,
		make([][]Edge, v),
	}
	return g
}

func (g *EdgeWeightedGraph) Vertices() int {
	return g.v
}

func (g *EdgeWeightedGraph) AddUndirectedEdge(e Edge) {
	v := e.From()
	w := e.OtherVertex(v)

	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
	g.e++
}

func (g *EdgeWeightedGraph) AddDirectedEdge(e Edge) {
	g.adj[e.From()] = append(g.adj[e.From()], e)

	g.e++
}

func (g *EdgeWeightedGraph) AdjacentTo(v int) []Edge {
	return g.adj[v]
}

func (g *EdgeWeightedGraph) Edges() []Edge {
	var b []Edge

	for v := 0; v < g.v; v++ {
		for _, e := range g.adj[v] {
			if e.OtherVertex(v) > v {
				b = append(b, e)
			}
		}
	}

	return b
}
