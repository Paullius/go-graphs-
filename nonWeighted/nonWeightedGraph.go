package nonWeighted

type NonWeightedGraph struct {
	v, e int
	adj  [][]int
}

func NewNonWeightedGraph(v int) NonWeightedGraph {
	g := NonWeightedGraph{
		v,
		0,
		make([][]int, v),
	}
	return g
}

// Number of vertices
func (g *NonWeightedGraph) Vertices() int {
	return g.v
}

// Number of edges
func (g *NonWeightedGraph) Edges() int {
	return g.e
}

// add edge to the graph
func (g *NonWeightedGraph) AddUndirectedEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *NonWeightedGraph) AddDirectedEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

// Vertices adjacent to v
func (g *NonWeightedGraph) AdjacentTo(v int) []int {
	return g.adj[v]
}

// Compute the degree of v
func (g *NonWeightedGraph) Degree(v int) int {
	return len(g.adj[v])
}

// Count of self-loops
func (g *NonWeightedGraph) NumberOfSelfLoops() int {
	count := 0
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			if v == w {
				count++
			}
		}
	}
	return count / 2
}

func (g *NonWeightedGraph) Reverse() NonWeightedGraph {
	dg := NewNonWeightedGraph(g.Vertices())
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			dg.AddDirectedEdge(w, v)
		}
	}
	return dg
}
