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

// when both goroutines are blocked you have a deadlock.

func main() {
	// create an unbuffered channel.
	ch := make(chan int)
	// create a goroutine.
	// without close(ch), the sending goroutine will block until it has more values to send.
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		// program will deadlock if we don't close the channel.
		close(ch)
	}()

	// range over channel to receive values.
	// the receiving goroutine will block until it receives another value or the ch is closed.
	for v := range ch {
		fmt.Printf("received: %v\n", v)
	}
	// when the channel closes the main goroutine closes.
}
