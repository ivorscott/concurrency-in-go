/*
This example shows interleaved out of order execution.
This occurs because the goroutines are running concurrently.

Output:
direct call
direct call
direct call
function call
wait for goroutines...
anonymous function
function value call
function call
anonymous function
function value call
function value call
anonymous function
function call
done...
*/

package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fun("direct call")

	go fun("function call")

	go func() {
		fun("anonymous function")
	}()

	routine := fun
	go routine("function value call")

	fmt.Println("wait for goroutines...")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done...")
}
