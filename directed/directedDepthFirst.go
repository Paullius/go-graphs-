package directed

type DirectedDepthFirst struct {
	g      *Digraph
	marked []bool
	count  int
}

func (g *Digraph) DirecteDepthFirst(s int) DirectedDepthFirst {
	df := DirectedDepthFirst{
		g,
		make([]bool, g.v),
		0,
	}
	df.search(s)
	return df
}

func (df *DirectedDepthFirst) search(v int) {
	df.marked[v] = true
	df.count++
	for _, w := range df.g.adj[v] {
		if !df.marked[w] {
			df.search(w)
		}
	}
}
