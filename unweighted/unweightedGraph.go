package unweighted

type UnweightedGraph struct {
	v, e int // v - count of vertices. e - count of edges
	adj  [][]int
}

func NewUnweightedGraph(v int) UnweightedGraph {
	g := UnweightedGraph{
		v,
		0,
		make([][]int, v),
	}
	return g
}

// Number of vertices
func (g *UnweightedGraph) Vertices() int {
	return g.v
}

// Number of edges
func (g *UnweightedGraph) Edges() int {
	return g.e
}

// add edge to the graph
func (g *UnweightedGraph) AddUndirectedEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *UnweightedGraph) AddDirectedEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

// Vertices adjacent to v
func (g *UnweightedGraph) AdjacentTo(v int) []int {
	return g.adj[v]
}

// Compute the degree of v
func (g *UnweightedGraph) Degree(v int) int {
	return len(g.adj[v])
}

// Count of self-loops
func (g *UnweightedGraph) NumberOfSelfLoops() int {
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

func (g *UnweightedGraph) Reverse() UnweightedGraph {
	dg := NewUnweightedGraph(g.Vertices())
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			dg.AddDirectedEdge(w, v)
		}
	}
	return dg
}
