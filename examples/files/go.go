package main

import (
	"fmt"
	"os"
)

func main() {
	// Writing to a file (basic)
	dataWrite := []byte("Hello, World!\nWritten from a Go program\n")
	err := os.WriteFile("example.txt", dataWrite, 666) // 666 -> Read/Write permissions for everyone
	if err != nil {
		panic(err)
	}

	// Reading a file (basic)
	dataRead, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Print(string(dataRead))
	}

}
