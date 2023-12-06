package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing filename argument")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("Attempting to the read the file: ", filename)
	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Printf("There was an error opening the file %v", filename)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
	fmt.Println()

}