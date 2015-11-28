package weighted

import (
	"container/heap"
	"github.com/paullius/go-graphs-/collections"
)

type LazyPrimMst struct {
	marked []bool
	mst    collections.Queue // MST edges
	pq     collections.PriorityQueue
}

func NewLazyPrimMst(g EdgeWeightedGraph) LazyPrimMst {
	l := LazyPrimMst{
		marked: make([]bool, g.v),
		mst:    collections.Queue{},
		pq:     collections.PriorityQueue{}}

	heap.Init(&l.pq)

	l.visit(g, 0)

	for l.pq.Len() > 0 {
		e := heap.Pop(&l.pq).(*Edge)
		v := e.AnyVertex()
		w := e.OtherVertex(v)
		if l.marked[v] && l.marked[w] {
			continue
		}

		l.mst.Push(e)
		if !l.marked[v] {
			l.visit(g, v)
		}
		if !l.marked[w] {
			l.visit(g, w)
		}
	}

	return l
}

func (l *LazyPrimMst) visit(g EdgeWeightedGraph, v int) {
	l.marked[v] = true

	for _, e := range g.AdjacentTo(v) {
		if l.marked[e.OtherVertex(v)] {
			heap.Push(&l.pq, e)
		}
	}

}

func (l *LazyPrimMst) Edges() []Edge {
	edges := make([]Edge, len(l.mst))
	for i, e := range l.mst {
		edges[i] = e.(Edge)
	}

	return edges
}
