/*
This example shows how to use sync/atomic on a counter
in a concurrently safe way.

Without the sync package, the variable counter is being accessed
by multiple goroutines that can be running in parallel and the
increment operation is not concurrent safe, so there will be a
data race and the data can get corrupted.

Output:
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
