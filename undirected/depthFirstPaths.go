package undirected

import (
	"fmt"
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

	return convert(path)
}

func convert(st collections.Stack) []int {
	b := make([]int, len(st))
	for i := range st {
		b[i] = st[i].(int)
	}
	return b
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

func (d *DepthFirstPaths) Print(s int) {
	fmt.Println("Depth-first Paths from ", s)

	for v := 0; v < d.g.v; v++ {

		fmt.Printf("%v to %v:", s, v)
		if d.HasPathTo(v) {
			for _, x := range d.PathTo(v) {
				if x == s {
					fmt.Printf("%v ", x)
				} else {
					fmt.Printf("-%v ", x)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
