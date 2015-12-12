package shortest

func DijkstrasExample(from int, to []int) {

	g := createShortestPathGraph()
	sp := NewDijkstraShortestPaths(g, from)

	sp.PrintShortestPath(to)
}

func createShortestPathGraph() EdgeWeightedDigraph {
	g := NewEdgeWeightedDigraph(8)
	g.AddEdge(NewEdge(4, 5, 0.35))
	g.AddEdge(NewEdge(5, 4, 0.35))
	g.AddEdge(NewEdge(4, 7, 0.37))
	g.AddEdge(NewEdge(5, 7, 0.28))
	g.AddEdge(NewEdge(7, 5, 0.28))
	g.AddEdge(NewEdge(5, 1, 0.32))
	g.AddEdge(NewEdge(0, 4, 0.38))
	g.AddEdge(NewEdge(0, 2, 0.26))
	g.AddEdge(NewEdge(7, 3, 0.39))
	g.AddEdge(NewEdge(1, 3, 0.29))
	g.AddEdge(NewEdge(2, 7, 0.34))
	g.AddEdge(NewEdge(6, 2, 0.40))
	g.AddEdge(NewEdge(3, 6, 0.52))
	g.AddEdge(NewEdge(6, 0, 0.58))
	g.AddEdge(NewEdge(6, 4, 0.93))

	return g
}
