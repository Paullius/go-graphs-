package weighted

import (
	"fmt"
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

	l.visit(g, 0)

	for !l.pq.IsEmpty() {
		e := l.pq.DelMin().(*Edge)
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
		if !l.marked[e.OtherVertex(v)] {
			l.pq.Insert(&e)
		}
	}

}

func (l *LazyPrimMst) Edges() []Edge {
	edges := make([]Edge, l.mst.Len())
	for i, e := range l.mst {
		edges[i] = *e.(*Edge)
	}

	return edges
}

func (l *LazyPrimMst) Weight() float32 {
	var weight float32 = 0.0
	for _, edge := range l.Edges() {
		weight += edge.weight
	}

	return weight
}

func (l *LazyPrimMst) PrintMinimumSpanningTree() {
	fmt.Println("Lazy version of Prim’s MST:")
	for _, e := range l.Edges() {
		fmt.Printf("%v", e)
	}

	fmt.Println()

	fmt.Println(l.Weight())
}
