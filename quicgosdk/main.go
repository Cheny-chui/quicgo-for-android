package main

import (
	"fmt"
	"quicgosdk/quicgosdk"
)

func main() {
	go func() { fmt.Print(quicgosdk.BuildServer("localhost:41420")) }()

	fmt.Print(quicgosdk.BuildClient("localhost:41420", "dddddd"))
}
