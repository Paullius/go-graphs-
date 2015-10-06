package undirected

type ConnectedComponents struct {
	g      *Graph
	marked []bool
	id     []int
	count  int
}

func (g *Graph) ConnectedComponents() ConnectedComponents {
	cc := ConnectedComponents{g, make([]bool, g.v), make([]int, g.v), 0}

	for s := 0; s <= cc.g.v; s++ {
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
	cc.count++
	for _, w := range cc.g.adj[v] {
		if !cc.marked[w] {
			cc.search(w)
		}
	}
}
