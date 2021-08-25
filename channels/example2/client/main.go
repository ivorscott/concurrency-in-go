package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// connect to server on localhost
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

// mustCopy copies the server response to std out.
func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal()
	}
}