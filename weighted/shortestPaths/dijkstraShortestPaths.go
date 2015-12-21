package shortestPaths

import (
	"fmt"
	"github.com/paullius/go-graphs-/collections"
	"github.com/paullius/go-graphs-/weighted"
	"math"
)

type DijkstraShortestPaths struct {
	edgeTo []weighted.Edge
	distTo []float32
	pq     collections.PriorityQueue
	from   int
}

func NewDijkstraShortestPaths(g weighted.EdgeWeightedGraph, from int) DijkstraShortestPaths {

	l := DijkstraShortestPaths{
		edgeTo: make([]weighted.Edge, g.Vertices()),
		distTo: make([]float32, g.Vertices()),
		pq:     collections.PriorityQueue{},
		from:   from}

	for v := 0; v < g.Vertices(); v++ { //initialize all vertice distance weight to maximum
		l.distTo[v] = math.MaxFloat32
	}
	l.distTo[from] = 0.0
	vertex := weighted.NewVertex(from, 0.0)
	l.pq.Insert(&vertex) //create zero index vertex with zero weight

	for l.pq.Len() > 0 {
		indexItem := l.pq.DelMin().(collections.IndexItem) // get lowest weight vertice from queue
		l.visit(g, indexItem.Index())
	}

	return l
}

func (sp *DijkstraShortestPaths) visit(g weighted.EdgeWeightedGraph, v int) {
	for _, e := range g.AdjacentTo(v) {
		var w = e.To() // to because directed graph
		totalWeighToVertex := sp.distTo[v] + e.Weight()
		if sp.distTo[w] > totalWeighToVertex { //is new total weigh to the vertex is lower
			sp.distTo[w] = totalWeighToVertex
			sp.edgeTo[w] = e
			if sp.pq.Contains(w) { //has lower weight than previous vertice
				sp.pq.Change(w, sp.distTo[w])
			} else { //there is no weight for the vertice
				vertex := weighted.NewVertex(w, sp.distTo[w])
				sp.pq.Insert(&vertex)
			}
		}
	}
}

func convert(st collections.Stack) []weighted.Edge {
	b := make([]weighted.Edge, len(st))
	for i := range st {
		b[i] = st[i].(weighted.Edge)
	}
	return b
}

func (sp *DijkstraShortestPaths) DistTo(v int) float32 {
	return sp.distTo[v]
}

func (sp *DijkstraShortestPaths) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat32
}

func (sp *DijkstraShortestPaths) PathTo(v int) []weighted.Edge {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := collections.Stack{}

	for e := sp.edgeTo[v]; e.Initialized(); e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	return convert(path)
}

func (sp *DijkstraShortestPaths) PrintShortestPath(to []int) {
	fmt.Println("Dijkstra’s Shortest Path:")

	for _, t := range to {
		fmt.Printf("%v to %v", sp.from, t)

		fmt.Printf(" (%v): ", sp.DistTo(t))
		if sp.HasPathTo(t) {
			for _, e := range sp.PathTo(t) {
				fmt.Printf("%v", e)
			}
		}
		fmt.Println()
	}
}
