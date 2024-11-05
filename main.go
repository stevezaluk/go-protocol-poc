package main

import (
	"fmt"

	"github.com/stevezaluk/go-protocol-poc/src"
)

func main() {
	fmt.Println("go protocol")

	serv := src.Server{}
	serv.Start()

	serv.Stop()
}
