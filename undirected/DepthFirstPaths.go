package undirected

type DepthFirstPaths struct {
	DF     DepthFirst
	EdgeTo []int
	S      int // source
}

func DepthFirstPathsSearch(g Graph, s int) DepthFirstPaths {
	var dfp DepthFirstPaths
	dfp.DF = DepthFirstSearch(g, s)
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
