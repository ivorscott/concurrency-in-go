/*
This example shows how properly encapsulate channel
responsibility in an owning function.

Output:
received: 0
received: 1
received: 2
received: 3
received: 4
done receiving
*/

package main

import "fmt"

func main() {

	// returns a receive only channel.
	owner := func() <-chan int {
		// owner creates the channel.
		ch := make(chan int)

		// defer close so that the channel gets closed at the end of the goroutine.
		go func() {
			defer close(ch) // owner closes channel when done.
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		// read values from channel.
		for v := range ch {
			fmt.Printf("received: %d\n", v)
		}
		fmt.Println("done receiving")
	}

	ch := owner()
	consumer(ch)
}
