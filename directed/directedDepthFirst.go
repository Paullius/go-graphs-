package directed

import (
	"github.com/paullius/go-graphs-/collections"
)

type DirectedDepthFirst struct {
	g      *Digraph
	marked []bool
	count  int
}

type DirectedCycle struct {
	cycle   collections.Stack
	onStack []bool
	edgeTo  []int
}

func (g *Digraph) DirecteDepthFirst(s int) DirectedDepthFirst {
	df := DirectedDepthFirst{
		g,
		make([]bool, g.v),
		0,
	}
	df.search(s, DirectedCycle{})
	return df
}

func (df *DirectedDepthFirst) search(v int, dc DirectedCycle) {
	dc.onStack[v] = true
	df.marked[v] = true
	df.count++
	for _, w := range df.g.adj[v] {
		if dc.hasCycle() {
			return
		}

		if !df.marked[w] {
			dc.edgeTo[w] = v
			df.search(w, dc)
		} else if dc.onStack[w] {
			dc.cycle = collections.Stack{}
			for x := v; x != w; x = dc.edgeTo[x] {
				dc.cycle.Push(x)
			}
			dc.cycle.Push(w)
			dc.cycle.Push(v)
		}
		dc.onStack[v] = false
	}
}

func (dc *DirectedCycle) hasCycle() bool {
	return dc.cycle != nil
}
