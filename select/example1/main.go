/*
This example shows using select to multiplex receiving on two channels.
Select enables us to process on the channels in the order in which the values arrive on the channel
rather than the order in which they are coded.

Output:
one
two
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// multiplex receive on channels
	// for loop required to receive both messages
	for i := 0; i < 2; i++ {
		select {
		case a := <-ch1:
			fmt.Println(a)
		case b := <-ch2:
			fmt.Println(b)
		}
	}
}
