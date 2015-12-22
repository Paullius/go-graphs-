package examples

import (
	"github.com/paullius/go-graphs-/weighted/minimumSpanningTree/examples"
	"github.com/paullius/go-graphs-/weighted/shortestPaths/examples"
	"sync"
)

func AsyncRun() {
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
