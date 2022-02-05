/*
This example shows how to use a buffered channel.

Output:
sending: 0
sending: 1
sending: 2
sending: 3
sending: 4
sending: 5
received: 0
received: 1
received: 2
received: 3
received: 4
received: 5
*/

package main

import "fmt"

// buffered channels only block when full.

func main() {
	// create a buffered channel.
	ch := make(chan int, 6)
	// create a goroutine.
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("sending: %d\n", i)
			// we can send up to 6 value without blocking.
			ch <- i
		}
		// program will deadlock if we don't close the channel.
		close(ch)
	}()

	// range over channel to receive values.
	for v := range ch {
		fmt.Printf("received: %d\n", v)
	}
}
