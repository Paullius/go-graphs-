package undirected

type Graph struct {
	V, E int
	adj  [][]int
}

func CreateGraph(v int) Graph {
	var g Graph
	g.V = v
	g.E = 0
	g.adj = make([][]int, v)
	return g
}

func (g *Graph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.E += g.E
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}
