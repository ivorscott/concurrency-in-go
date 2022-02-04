/*
The output in this example shows a client/server relationship.
The server handles concurrent connections to multiple clients.

Open 3 terminal windows. 2 for 2 clients, and 1 for the server.

Example output:
	./server

	./client
	response from server
	response from server
	response from server

	./client
	response from server
	response from server
	response from server
*/
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// write server program to handle concurrent client connections
	listener, err := net.Listen("tcp", "localhost:8000")
	if err !=nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// continue to next connection
			continue
		}
		go handleConn(conn)
	}
}

// handleConn writes a response to the client connection.
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c,"response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}