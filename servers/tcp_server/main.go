package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080") // Network to listen, port to listen
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		io.WriteString(conn, "Hello")
		fmt.Fprintln(conn, "ULULU")
		fmt.Fprintf(conn, "%v", "So funny")
	}
}
