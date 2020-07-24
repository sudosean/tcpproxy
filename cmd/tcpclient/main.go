package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "seanlarge.me:80")
	if err != nil {
		log.Fatalln("Unable to connect")
	}
	defer dst.Close()

	// run in go routing to prevent io.Copy from blocking
	go func() {
		// copy source output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// copy out destination output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main(){
	// listen on local port 80
	listener, err := net.Listen("tcp", "80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
