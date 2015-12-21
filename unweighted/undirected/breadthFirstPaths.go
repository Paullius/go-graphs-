package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/collections"
	"github.com/paullius/go-graphs-/unweighted"
)

type BreadthFirstPaths struct {
	g      unweighted.NonWeightedGraph
	marked []bool
	edgeTo []int
	s      int
}

func NewBreadthFirstPaths(g unweighted.NonWeightedGraph, s int) BreadthFirstPaths {
	bfp := BreadthFirstPaths{
		g,
		make([]bool, g.Vertices()),
		make([]int, g.Vertices()),
		s,
	}

	bfp.search(s)
	return bfp
}

// Is there a path from s to v
func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	return bfp.marked[v]
}

func (bfp *BreadthFirstPaths) PathTo(v int) []int {
	if !bfp.HasPathTo(v) {
		return []int{}
	}
	path := collections.Stack{}
	for x := v; x != bfp.s; x = bfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(bfp.s)

	return path.ConvertToInt()
}

func (bfp *BreadthFirstPaths) search(s int) {
	queue := collections.Queue{}
	bfp.marked[s] = true
	queue.Push(s)

	for queue.Len() > 0 {
		v := queue.Pop().(int)
		for _, w := range bfp.g.AdjacentTo(v) {
			if !bfp.marked[w] {
				bfp.edgeTo[w] = v
				bfp.marked[w] = true
				queue.Push(w)
			}
		}
	}
}

func (bfp *BreadthFirstPaths) Print(s int) {
	fmt.Println("Breadth-first Paths from ", s)

	for v := 0; v < bfp.g.Vertices(); v++ {

		fmt.Printf("%v to %v:", s, v)
		if bfp.HasPathTo(v) {
			for _, x := range bfp.PathTo(v) {
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
