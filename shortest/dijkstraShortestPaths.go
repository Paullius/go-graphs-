package shortest

import (
	"container/heap"
	"github.com/paullius/go-graphs-/collections"
	"math"
)

type DijkstraShortestPaths struct {
	edgeTo []DirectedEdge
	distTo []float32
	pq     collections.IndexMinPriorityQueue
}

func NewDijkstraShortestPaths(g EdgeWeightedDigraph, s int) DijkstraShortestPaths {

	l := DijkstraShortestPaths{
		edgeTo: make([]DirectedEdge, g.v),
		distTo: make([]float32, g.v),
		pq:     collections.IndexMinPriorityQueue{}}

	heap.Init(&l.pq)

	for v := 0; v < g.v; v++ {
		distTo[v] = math.Inf(1)
	}
	distTo[s] = 0.0

	heap.Push(&l.pq, 0.0)

	for l.pq.Len() > 0 {
		e := heap.Pop(&l.pq).(*float32)

	}
}

func (sp *DijkstraShortestPaths) relax( g EdgeWeightedDigraph,  v int)
{
	for _, e := range g.adj(v) {
	{
		var w = e.to();
		if distTo[w] > distTo[v] + e.weight()
		{
			sp.distTo[w] = sp.distTo[v] + e.weight()
			sp.edgeTo[w] = e
		if sp.pq.Contains(w) {
			sp.pq.Change(w, sp.distTo[w])
		}
		else {
			sp.pq.insert(w, sp.distTo[w])
		}
	}
}

func convert(st collections.Stack) []DirectedEdge {
	b := make([]DirectedEdge, len(st))
	for i := range st {
		b[i] = st[i].(DirectedEdge)
	}
	return b
}

func (sp *DijkstraShortestPaths) DistTo(v int) float32 {
   return sp.distTo[v]
}

func (sp *DijkstraShortestPaths) HasPathTo(v int) bool {
   return sp.distTo[v] < math.Inf(1)
}

func (sp *DijkstraShortestPaths) PathTo(v int) []DirectedEdge {
   if (!sp.hasPathTo(v)) return nil

    path := collections.Stack{}
     for e := edgeTo[v]; e != null; e = edgeTo[e.from()] {
     	  path.push(e)
     }
     return convert(path) 
}

