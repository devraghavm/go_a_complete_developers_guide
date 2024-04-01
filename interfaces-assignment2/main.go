package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	readFile(filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", filename)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
