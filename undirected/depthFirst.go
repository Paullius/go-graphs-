package undirected

import (
	"fmt"
)

type DepthFirst struct {
	g      *Graph
	marked []bool
	count  int
}

// Do depth first graph search
// s - source vertex
func (g *Graph) DepthFirst(s int) DepthFirst {
	df := DepthFirst{
		g,
		make([]bool, g.v),
		0,
	}
	df.search(s)
	return df
}

// Count of vertices connected to s
func (df *DepthFirst) ConnectedCount() int {
	return df.count
}

// Is v connected to s
func (df *DepthFirst) IsConnected(v int) bool {
	return df.marked[v]
}

func (df *DepthFirst) search(v int) {
	df.marked[v] = true
	df.count++
	for _, w := range df.g.adj[v] {
		if !df.marked[w] {
			df.search(w)
		}
	}
}

func (d *DepthFirst) Print(s int) {
	fmt.Println("Depth-first from ", s)

	for v := 0; v < d.g.v; v++ {
		if d.marked[v] {
			fmt.Printf("%v ", v)
		}

	}
	fmt.Println()

	if d.count != d.g.v {
		fmt.Printf("NOT ")
	}
	fmt.Println("connected")
}
