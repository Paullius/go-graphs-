package undirected

type DepthFirst struct {
  G Graph
  Marked [] bool
  Count int
}

//s - source vertex
func DepthFirstSearch (g Graph, s int) DepthFirst {
   var df DepthFirst
   df.G = g
   df.Marked = make([]bool,g.V)
   depthFirstSearch(df,s)

   return df;
}

//is v connected to s
func depthFirstSearch (df DepthFirst, v int) {
  df.Marked[v] = true
  df.Count++;

  for _,w := range df.G.adj[v] {
     if(!df.Marked[w]) {
     	depthFirstSearch(df,w)
     }
  }
}