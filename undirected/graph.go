 
package graphundirected

type Graph struct {
	v,e int
	adj [][]int
}

func create(v int) Graph {
  var g Graph
  g.v =v
  g.e = 0
  g.adj = make([][]int,v)
  return g
}

func addEdge(v int, w int, g Graph) {
	g.adj[v] = append(g.adj[v],w)
	g.adj[w] = append(g.adj[w],v)
	g.e += g.e
}

func adj(v int, g Graph) []int {
  return g.adj[v]
}
