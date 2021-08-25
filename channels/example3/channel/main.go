/*
This example shows how to get a values from a goroutine.
In Go we share memory by communicating (with channels).

Output:
computed value 3
*/

package main

import "fmt"

func main() {
	// create a channel
	ch := make(chan int)
	// create a goroutine
	go func(a,b int) {
		c := a+b
		ch <- c // send the value
	}(1,2)

	c := <-ch // retrieve the value
	fmt.Printf("computed value %v\n", c)
}

