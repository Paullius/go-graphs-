package undirected

type DepthFirstPaths struct {
	DF     DepthFirst
	EdgeTo []int
	S      int // source
}

func (g *Graph) DepthFirstPathsSearch(s int) DepthFirstPaths {
	var dfp DepthFirstPaths
	dfp.DF = g.DepthFirstSearch(s)
	dfp.EdgeTo = make([]int, g.V)
	dfp.S = s

	depthFirstPathsSearch(dfp, s)

	return dfp
}

func depthFirstPathsSearch(dfp DepthFirstPaths, v int) {
	dfp.DF.Marked[v] = true
	dfp.DF.Count++

	for _, w := range dfp.DF.G.adj[v] {
		if !dfp.DF.Marked[w] {
			dfp.EdgeTo[w] = v
			depthFirstPathsSearch(dfp, w)
		}
	}
}
