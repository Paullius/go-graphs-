package undirected

type DepthFirstPaths struct {
	DF     DepthFirst
	EdgeTo []int
	S      int // source
}

func (g *Graph) DepthFirstPaths(s int) DepthFirstPaths {
	dfp := DepthFirstPaths{
		g.DepthFirst(s),
		make([]int, g.v),
		s,
	}
	dfp.search(s)
	return dfp
}

func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	return dfp.DF.Marked[v]
}

func (dfp *DepthFirstPaths) PathTo(v int) []int {
	if !dfp.HasPathTo(v) {
		return []int{}
	}
	path := Stack{}
	for x := v; x != dfp.S; x = dfp.EdgeTo[x] {
		path.Push(x)
	}
	path.Push(dfp.S)
	return path
}

func (dfp *DepthFirstPaths) search(v int) {
	dfp.DF.Marked[v] = true
	dfp.DF.Count++
	for _, w := range dfp.DF.G.adj[v] {
		if !dfp.DF.Marked[w] {
			dfp.EdgeTo[w] = v
			dfp.search(w)
		}
	}
}
