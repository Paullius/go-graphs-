package undirected

type DepthFirstPaths struct {
	df     DepthFirst
	edgeTo []int
	s      int // source
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

func (dfp *DepthFirstPaths) DF() DepthFirst {
	return dfp.df
}

// Is there a path from s to v
//TODO: move to interface
func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	return dfp.df.marked[v]
}

// Path from s to v
func (dfp *DepthFirstPaths) PathTo(v int) []int {
	if !dfp.HasPathTo(v) {
		return []int{}
	}
	path := Stack{}
	for x := v; x != dfp.s; x = dfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(dfp.s)
	return path
}

func (dfp *DepthFirstPaths) search(v int) {
	dfp.df.marked[v] = true
	dfp.df.count++
	for _, w := range dfp.df.g.adj[v] {
		if !dfp.df.marked[w] {
			dfp.edgeTo[w] = v
			dfp.search(w)
		}
	}
}
