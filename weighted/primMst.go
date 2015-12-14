package weighted

import (
	"fmt"
	"github.com/paullius/go-graphs-/collections"
	"math"
)

type PrimMst struct {
	marked []bool
	edgeTo []Edge
	distTo []float32
	pq     collections.IndexMinPriorityQueue
}

func NewPrimMst(g EdgeWeightedGraph) PrimMst {
	l := PrimMst{
		marked: make([]bool, g.v),
		edgeTo: make([]Edge, g.v),
		distTo: make([]float32, g.v),
		pq:     collections.IndexMinPriorityQueue{}}

	for v := 0; v < g.v; v++ {
		l.distTo[v] = math.MaxFloat32
	}
	l.distTo[0] = 0.0
	l.pq.Insert(0, 0.0)

	for !l.pq.IsEmpty() {
		l.visit(g, l.pq.DelMin())
	}

	return l
}

func (l *PrimMst) visit(g EdgeWeightedGraph, v int) {
	l.marked[v] = true

	for _, e := range g.AdjacentTo(v) {
		w := e.OtherVertex(v)
		if l.marked[w] {
			continue
		}
		if e.Weight() < l.distTo[w] {
			l.edgeTo[w] = e
			l.distTo[w] = e.Weight()
			if l.pq.Contains(w) {
				l.pq.Change(w, l.distTo[w])
			} else {
				l.pq.Insert(w, l.distTo[w])
			}
		}
	}

}

func (l *PrimMst) Edges() []Edge {

	return l.edgeTo
}

func (l *PrimMst) Weight() float32 {
	var weight float32 = 0.0
	for _, edge := range l.Edges() {
		weight += edge.weight
	}

	return weight
}

func (l *PrimMst) PrintMinimumSpanningTree() {
	fmt.Println("Prim’s MST:")
	for _, e := range l.Edges() {
		fmt.Printf("%v", e)
		fmt.Println()
	}

	fmt.Println("Weight: ", l.Weight())
}
