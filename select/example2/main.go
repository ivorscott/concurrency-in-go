/*
This example shows using select with a timeout duration.

Output when time.After is never reached:
one

Output when time.After is reached:
timeout
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(1 * time.Second): // try increasing time after value
		fmt.Println("timeout")
	}
}
