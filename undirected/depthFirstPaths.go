package undirected

import (
	"github.com/paullius/go-graphs-/collections"
)

type DepthFirstPaths struct {
	g      *Graph
	marked []bool
	edgeTo []int
	s      int // source
	count  int
}

func (g *Graph) DepthFirstPaths(s int) DepthFirstPaths {
	dfp := DepthFirstPaths{
		g,
		make([]bool, g.v),
		make([]int, g.v),
		s,
		0,
	}
	dfp.search(s)
	return dfp
}

// Count of vertices connected to s
func (dfp *DepthFirstPaths) ConnectedCount() int {
	return dfp.count
}

// Is there a path from s to v
//TODO: move to interface
func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	return dfp.marked[v]
}

// Path from s to v
func (dfp *DepthFirstPaths) PathTo(v int) []int {
	if !dfp.HasPathTo(v) {
		return []int{}
	}
	path := collections.Stack{}
	for x := v; x != dfp.s; x = dfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(dfp.s)
	return path
}

func (dfp *DepthFirstPaths) search(v int) {
	dfp.marked[v] = true
	dfp.count++
	for _, w := range dfp.g.adj[v] {
		if !dfp.marked[w] {
			dfp.edgeTo[w] = v
			dfp.search(w)
		}
	}
}
