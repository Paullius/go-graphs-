package minimumSpanningTree

import (
	"fmt"
	"github.com/paullius/go-graphs-/collections"
	"github.com/paullius/go-graphs-/weighted"
	"math"
)

type PrimMst struct {
	marked []bool
	edgeTo []weighted.Edge
	distTo []float32
	pq     collections.PriorityQueue
}

func NewPrimMst(g weighted.EdgeWeightedGraph) PrimMst {
	l := PrimMst{
		marked: make([]bool, g.Vertices()),
		edgeTo: make([]weighted.Edge, g.Vertices()),
		distTo: make([]float32, g.Vertices()),
		pq:     collections.PriorityQueue{}}

	for v := 0; v < g.Vertices(); v++ {
		l.distTo[v] = math.MaxFloat32
	}
	l.distTo[0] = 0.0
	l.pq.Insert(weighted.NewVertex(0, 0.0))

	for !l.pq.IsEmpty() {
		indexItem := l.pq.DelMin().(collections.IndexItem)
		l.visit(g, indexItem.Index())
	}

	return l
}

func (l *PrimMst) visit(g weighted.EdgeWeightedGraph, v int) {
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
				l.pq.Insert(weighted.NewVertex(w, l.distTo[w]))
			}
		}
	}

}

func (l *PrimMst) Edges() []weighted.Edge {

	return l.edgeTo
}

func (l *PrimMst) Weight() float32 {
	var weight float32 = 0.0
	for _, edge := range l.Edges() {
		weight += edge.Weight()
	}

	return weight
}

func (l *PrimMst) PrintMinimumSpanningTree() {
	fmt.Println("Prim’s Minimum Spanning Tree:")
	for _, e := range l.Edges() {
		fmt.Printf("%v", e)
		fmt.Println()
	}

	fmt.Println("Weight: ", l.Weight())
}
