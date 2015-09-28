package undirected

type BreadthFirstPaths struct {
	g      *Graph
	marked []bool
	edgeTo []int
	s      int
}

func (g *Graph) BreadthFirstPaths(s int) BreadthFirstPaths {
	bfp := BreadthFirstPaths{
		g,
		make([]bool, g.v),
		make([]int, g.v),
		s,
	}

	bfp.search(s)
	return bfp
}

// Is there a path from s to v
//TODO: move to interface
func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	return bfp.marked[v]
}

func (bfp *BreadthFirstPaths) search(s int) {
	queue := Queue{}
	bfp.marked[s] = true
	queue.Push(s)

	for queue.Len() > 0 {
		var v = queue.Pop()
		for _, w := range bfp.g.adj[v] {
			if !bfp.marked[w] {
				bfp.edgeTo[w] = v
				bfp.marked[w] = true
				queue.Push(w)
			}
		}
	}
}
