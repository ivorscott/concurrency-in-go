/*
This example shows using the sync package to share a shared resource
between two goroutines.

The deposit and the withdrawal operation must be atomic.

If goroutines are running concurrently in interleaved manner,
preempting each other, then data race is possible and data can
get corrupted.

We need to use a mutex lock to guard access.

Output:
0
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var lk sync.Mutex

	// balance is a shared resource
	deposit := func(amount int) {
		lk.Lock()
		balance += amount
		lk.Unlock()
	}
	withdrawal := func(amount int) {
		lk.Lock()
		defer lk.Unlock()
		balance -= amount
	}

	// make 100 deposits of $1
	// and 100 withdrawals of $1 concurrently
	// run the program and check result

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}
