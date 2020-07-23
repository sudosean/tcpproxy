package main

import (
	"fmt"
	"io"
	"log"
	"net"
)
// echo is a handler func that echoes data
func echo(conn net.Conn)  {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		// read data
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		log.Println("Writing data")
		// send data
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}

	}
}

func main() {
	fmt.Println("Hello from echo server")

	// tcp listner on port 20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Cannot bind port")
	}
	log.Println("listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		log.Println("Received connection")
		go echo(conn)
	}
}
