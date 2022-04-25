package main

import (
	"fmt"
	"quicgosdk/quicgo"
)

/*
	for test the lib in go
*/
func main() {
	go func() { fmt.Print(quicgo.BuildServer("localhost:41420")) }()

	fmt.Print(quicgo.BuildClient("localhost:41420", "dddddd"))
}
