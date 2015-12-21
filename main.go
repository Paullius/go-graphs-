package main

import (
	"fmt"
	"github.com/paullius/go-graphs-/nonWeighted/directed/examples"
	"github.com/paullius/go-graphs-/nonWeighted/undirected/examples"
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree/examples"
	"github.com/paullius/go-graphs-/weighted/shortestPaths/examples"
	"log"
	"net/http"
	"sync"
)

func main() {

	undirected.DepthFirstExample(0)
	undirected.ConnectedComponentsExample()
	undirected.DepthFirstPathsExample(0)
	undirected.BreadthFirstPathsExample(0)

	directed.DirectedDepthFirstExample(2)

	minimumSpanningTree.PrimMstExample()

	shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})

	asyncRun()

	startWebServer()

}

func asyncRun() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		minimumSpanningTree.PrimMstExample()
		defer wg.Done()
	}()
	go func() {
		shortestPaths.DijkstrasExample(0, []int{1, 2, 3, 4, 5, 6, 7})
		defer wg.Done()
	}()
	wg.Wait()
}

func startWebServer() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Server on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "<html>Hello, <b>User</b>:)<html>")
}

//go run %GOPATH%\src\github.com\Paullius\go-graphs-\main.go
//go install github.com/paullius/go-graphs-
//go test github.com/paullius/go-graphs-\nonWeighted
