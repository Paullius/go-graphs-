package shortest

import (
	"fmt"
	"github.com/paullius/go-graphs-/collections"
	"math"
)

type DijkstraShortestPaths struct {
	edgeTo []DirectedEdge
	distTo []float32
	pq     collections.IndexMinPriorityQueue
	from   int
}

func NewDijkstraShortestPaths(g EdgeWeightedDigraph, from int) DijkstraShortestPaths {

	l := DijkstraShortestPaths{
		edgeTo: make([]DirectedEdge, g.v),
		distTo: make([]float32, g.v),
		pq:     collections.IndexMinPriorityQueue{},
		from:   from}

	for v := 0; v < g.v; v++ {
		l.distTo[v] = math.MaxFloat32
	}
	l.distTo[from] = 0.0

	l.pq.Insert(from, 0.0)

	for l.pq.Len() > 0 {
		l.relax(g, l.pq.DelMin())
	}

	return l
}

func (sp *DijkstraShortestPaths) relax(g EdgeWeightedDigraph, v int) {
	for _, e := range g.AdjacentTo(v) {
		var w = e.To()
		if sp.distTo[w] > sp.distTo[v]+e.Weight() {
			sp.distTo[w] = sp.distTo[v] + e.Weight()
			sp.edgeTo[w] = e
			if sp.pq.Contains(w) {
				sp.pq.Change(w, sp.distTo[w])
			} else {
				sp.pq.Insert(w, sp.distTo[w])
			}
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
	return sp.distTo[v] < math.MaxFloat32
}

func (sp *DijkstraShortestPaths) PathTo(v int) []DirectedEdge {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := collections.Stack{}

	for e := sp.edgeTo[v]; !(e.From() == 0 && e.To() == 0); e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	// l := len(sp.edgeTo)
	// i := v
	// for i < l {
	// 	e := sp.edgeTo[i]
	// 	path.Push(e)
	// 	i = e.From()
	// }

	return convert(path)
}

func (sp *DijkstraShortestPaths) PrintShortestPath(to int) {
	fmt.Println("Dijkstra’s Shortest Path:")
	fmt.Printf("%v to %v", sp.from, to)

	fmt.Printf(" (%v): ", sp.DistTo(to))
	if sp.HasPathTo(to) {
		for _, e := range sp.PathTo(to) {
			fmt.Printf("%v", e)
		}
	}
	fmt.Println()
}
