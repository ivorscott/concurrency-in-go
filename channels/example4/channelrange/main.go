/*
This example shows how to for range over values from a goroutine.

Output:
received: 0
received: 1
received: 2
received: 3
received: 4
received: 5
*/

package main

import "fmt"

func main() {
	// create a channel
	ch := make(chan int)
	// create a goroutine
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		// program will deadlock if we don't close the channel
		close(ch)
	}()

	// range over channel to receive values
	for v := range ch {
		fmt.Printf("received: %v\n", v)
	}
}
