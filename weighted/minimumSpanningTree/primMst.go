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

	for v := 0; v < g.Vertices(); v++ { //initialize all vertice distance weight to maximum
		l.distTo[v] = math.MaxFloat32
	}
	l.distTo[0] = 0.0
	vertex := weighted.NewVertex(0, 0.0)
	l.pq.Insert(&vertex) //create zero index vertex with zero weight

	for !l.pq.IsEmpty() {
		vertex := l.pq.DelMin().(collections.IndexItem) // get lowest weight vertice from queue
		l.visit(g, vertex.Index())
	}

	return l
}

func (l *PrimMst) visit(g weighted.EdgeWeightedGraph, v int) {
	l.marked[v] = true

	//finds not visited lowest weight vertice
	for _, e := range g.AdjacentTo(v) {
		w := e.OtherVertex(v) //get vertex connected with the v vertex
		if l.marked[w] {      //vertice was already visited
			continue
		}
		if e.Weight() < l.distTo[w] {
			l.edgeTo[w] = e
			l.distTo[w] = e.Weight()
			if l.pq.Contains(w) { //has lower weight than previous vertice
				l.pq.Change(w, l.distTo[w])
			} else { //there is no weight for the vertice
				vertex := weighted.NewVertex(w, l.distTo[w])
				l.pq.Insert(&vertex)
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
