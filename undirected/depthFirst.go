package undirected

type DepthFirst struct {
	G      *Graph
	Marked []bool
	Count  int
}

//s - source vertex
func (g *Graph) DepthFirst(s int) DepthFirst {
	df := DepthFirst{
		g,
		make([]bool, g.v),
		0,
	}
	df.search(s)
	return df
}

//is v connected to s
func (df *DepthFirst) search(v int) {
	df.Marked[v] = true
	df.Count++
	for _, w := range df.G.adj[v] {
		if !df.Marked[w] {
			df.search(w)
		}
	}
}
