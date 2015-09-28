package undirected

type Graph struct {
	v, e int
	adj  [][]int
}

func NewGraph(v int) Graph {
	g := Graph{
		v,
		0,
		make([][]int, v),
	}
	return g
}

func (g *Graph) Vertices() int {
	return g.v
}

func (g *Graph) Edges() int {
	return g.e
}

func (g *Graph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e += g.e
}

func (g *Graph) AdjacentTo(v int) []int {
	return g.adj[v]
}

func (g *Graph) Degree(v int) int {
	return len(g.adj[v])
}

func (g *Graph) NumberOfSelfLoops() int {
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
