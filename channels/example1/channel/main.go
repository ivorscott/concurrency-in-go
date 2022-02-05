/*
This example shows how to get a values from a goroutine.
In Go we share memory by communicating (with channels).

Output:
computed value 3
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutines are asynchronous.

// the main function is the main goroutine.
func main() {
	// create a unbuffered channel.
	ch := make(chan int)
	// create a goroutine with an anonymous function.
	// this operates as the sending goroutine.
	// the channel, "ch", is available within the child goroutine because
	// it exists within the parent context.
	go func(a, b int) {
		c := a + b
		ch <- c // send the value to the channel. this blocks until a receiver goroutine receives from the channel.
		fmt.Println("done")
	}(1, 2)
	fmt.Printf("%d goroutines exist at this point.\n", runtime.NumGoroutine())

	c := <-ch // retrieve the value.
	fmt.Printf("computed value %v\n", c)
	time.Sleep(1 * time.Second) // without this the main goroutine finishes before "done" is printed.

	// When the anonymous goroutine finishes executing it gets reclaimed by the
	fmt.Printf("only the %d goroutine remains: the main goroutine. \n", runtime.NumGoroutine())

}
