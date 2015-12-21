package directed

import (
	"github.com/paullius/go-graphs-/collections"
	"github.com/paullius/go-graphs-/unweighted"
)

type DirectedDepthFirstWithCycleDetection struct {
	cycle              collections.Stack
	onStack            []bool
	edgeTo             []int
	directedDepthFirst DirectedDepthFirst
}

func NewDirectedDepthFirstWithCycleDetection(g unweighted.UnweightedGraph, s int) DirectedDepthFirst {

	cd := DirectedDepthFirstWithCycleDetection{directedDepthFirst: DirectedDepthFirst{
		g,
		make([]bool, g.Vertices()),
		0,
	}}

	cd.search(s)
	return cd.directedDepthFirst
}

func (df *DirectedDepthFirstWithCycleDetection) search(v int) {
	df.onStack[v] = true
	df.directedDepthFirst.marked[v] = true
	df.directedDepthFirst.count++
	for _, w := range df.directedDepthFirst.g.AdjacentTo(v) {
		if df.hasCycle() {
			return
		}

		if !df.directedDepthFirst.marked[w] {
			df.edgeTo[w] = v
			df.search(w)
		} else if df.onStack[w] {
			df.cycle = collections.Stack{}
			for x := v; x != w; x = df.edgeTo[x] {
				df.cycle.Push(x)
			}
			df.cycle.Push(w)
			df.cycle.Push(v)
		}
		df.onStack[v] = false
	}
}

func (dc *DirectedDepthFirstWithCycleDetection) hasCycle() bool {
	return dc.cycle != nil
}
