package undirected

type BreadthFirstPaths struct {
	G      *Graph
	Marked []bool
	EdgeTo []int
	S      int
}

func (g *Graph) BreadthFirstPaths(s int) BreadthFirstPaths {
	bfp := BreadthFirstPaths{
		g,
		make([]bool, g.v),
		make([]int, g.v),
		s,
	}

	bfp.breadthFirstSearch(s)
	return bfp
}

func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	return bfp.Marked[v]
}

func (bfp *BreadthFirstPaths) breadthFirstSearch(s int) {
	queue := Queue{}
	bfp.Marked[s] = true
	queue.Push(s)

	for queue.Len() > 0 {
		var v = queue.Pop()
		for _, w := range bfp.G.adj[v] {
			if !bfp.Marked[w] {
				bfp.EdgeTo[w] = v
				bfp.Marked[w] = true
				queue.Push(w)
			}
		}
	}
}
