// Proxy mediates TCP traffic between client and upstream. Adapted from
// youtu.be/J4J-A9tcjcA.
//
// Level: intermediate
// Topics: net, security, concurrency
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for { // accept loop
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// NOTE: don't put any blocking code here!
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close() // to release precious file descriptor

	upstream, err := net.Dial("tcp", "google.com:http")
	if err != nil {
		log.Print(err)
		return
	}
	defer upstream.Close()

	go io.Copy(upstream, conn) // in this case it's ok not track goroutine
	io.Copy(conn, upstream)
}
