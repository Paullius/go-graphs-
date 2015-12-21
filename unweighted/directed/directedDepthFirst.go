package directed

import (
	"fmt"
	"github.com/paullius/go-graphs-/unweighted"
)

type DirectedDepthFirst struct {
	g      unweighted.UnweightedGraph
	marked []bool
	count  int
}

func NewDirecteDepthFirst(g unweighted.UnweightedGraph, s int) DirectedDepthFirst {
	df := DirectedDepthFirst{
		g,
		make([]bool, g.Vertices()),
		0,
	}
	df.search(s)
	return df
}

func (df *DirectedDepthFirst) search(v int) {
	df.marked[v] = true
	df.count++
	for _, w := range df.g.AdjacentTo(v) {
		if !df.marked[w] {
			df.search(w)
		}
	}
}

func (df *DirectedDepthFirst) Print(s int) {
	fmt.Println("Directed Depth-first Paths from ", s)

	for v := 0; v < df.g.Vertices(); v++ {
		if df.marked[v] {
			fmt.Printf("%v ", v)
		}
	}
	fmt.Println()
}
