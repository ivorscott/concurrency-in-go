/*
This example shows using select with non-blocking communication.

Output:
*/
package main

import (
	"fmt"
	"time"
)

// Blocks
//func main() {
//	ch := make(chan string)
//
//	go func() {
//		for i := 0; i < 3; i++ {
//			time.Sleep(1 * time.Second)
//			ch <- "message"
//		}
//	}()
//
//	// if there is no value on channel, do not block
//	for i := 0; i < 2; i++ {
//		m := <-ch
//		fmt.Println(m)
//	}

// Do some processing
//  fmt.Println("processing")
//	time.Sleep(1500 * time.Millisecond)
//}

// Non-blocking
func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	// if there is no value on channel, do not block
	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no message received")
		}
		// Do some processing
		fmt.Println("processing")
		time.Sleep(1500 * time.Millisecond)
	}
}
