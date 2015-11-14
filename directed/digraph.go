package directed

type Digraph struct {
	v, e int
	adj  [][]int
}

// Creates graph
func NewDigraph(v int) Digraph {
	g := Digraph{
		v,
		0,
		make([][]int, v),
	}
	return g
}

// Number of vertices
func (g *Digraph) Vertices() int {
	return g.v
}

// Number of edges
func (g *Digraph) Edges() int {
	return g.e
}

// add edge to the graph
func (g *Digraph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

func (g *Digraph) Reverse() Digraph {
	dg := NewDigraph(g.Vertices())
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			dg.AddEdge(w, v)
		}
	}
	return dg
}
