 
package undirected

type Graph struct {
	V,E int
	adj [][]int
}

func Create(v int) Graph {
  var g Graph
  g.V =v
  g.E = 0
  g.adj = make([][]int,v)
  return g
}

func AddEdge(v int, w int, g Graph) {
	g.adj[v] = append(g.adj[v],w)
	g.adj[w] = append(g.adj[w],v)
	g.E += g.E
}

func Adj(v int, g Graph) []int {
  return g.adj[v]
}
