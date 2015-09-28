package undirected

type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n int) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

type BreadthFirstPaths struct {
	G      *Graph
	Marked []bool
	EdgeTo []int
	S      int
}

func (g *Graph) BreadthFirstPathsSearch(s int) BreadthFirstPaths {
	var bfp BreadthFirstPaths
	bfp.G = g
	bfp.Marked = make([]bool, g.V)
	bfp.EdgeTo = make([]int, g.V)
	bfp.S = s

	breadthFirstSearch(bfp, s)
	return bfp
}

func (bfp *BreadthFirstPaths) HasPathTo(v int) bool {
	return bfp.Marked[v]
}

func breadthFirstSearch(bfp BreadthFirstPaths, s int) {
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
