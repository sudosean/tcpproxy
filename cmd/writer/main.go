package main

import (
	"fmt"
	"os"
)
type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

type FooWriter struct{}
func(fooWriter *FooWriter) Write(b []byte) (int,error) {
	fmt.Print("out <")
	return os.Stdout.Write(b)
}

func main()  {
	fmt.Println("Hello from writter")
	var (
		reader FooReader
	)
}
