package undirected

import (
	"fmt"
	"github.com/paullius/go-graphs-/nonWeighted"
)

type ConnectedComponents struct {
	g      nonWeighted.NonWeightedGraph
	marked []bool
	id     []int
	count  int
}

func NewConnectedComponents(g nonWeighted.NonWeightedGraph) ConnectedComponents {
	cc := ConnectedComponents{g, make([]bool, g.Vertices()), make([]int, g.Vertices()), 0}

	for s := 0; s < cc.g.Vertices(); s++ {
		if !cc.marked[s] {
			cc.search(s)
			cc.count++
		}
	}

	return cc
}

func (cc *ConnectedComponents) Connected(v int, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *ConnectedComponents) Count() int {
	return cc.count
}

func (cc *ConnectedComponents) search(v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range cc.g.AdjacentTo(v) {
		if !cc.marked[w] {
			cc.search(w)
		}
	}
}

func (cc *ConnectedComponents) Print() {
	fmt.Println("Depth-first search to find connected components: ", cc.count)

	components := make([][]int, cc.count)

	for v := 0; v < cc.g.Vertices(); v++ {
		components[cc.id[v]] = append(components[cc.id[v]], v)
	}

	for _, i := range components {
		for _, v := range i {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
	fmt.Println()

}
