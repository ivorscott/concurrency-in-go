/*
This example shows how to implement relaying of messages with
channel direction.

Output:
message
*/

package main

import "fmt"

// constrain functions in terms of what they can do with channels.
// use channel direction to describe send only or receive only channels.

// genMsg expects a send only channel.
// If we try to receive from the channel we get a compiler error.
func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "message"
}

// relayMsh expects 1 receive only channel and 1 send only channel.
func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// receive message on ch1.
	// send message on ch2.
	m := <-ch1
	ch2 <- m
}

func main() {
	// create ch1 and ch2.
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	v := <-ch2
	fmt.Println(v)
}
