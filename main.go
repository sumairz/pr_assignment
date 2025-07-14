package main

import (
	"fmt"
	"fsm/modthree"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go <binary_string>\n")
		fmt.Fprintf(os.Stderr, "Example: go run main.go 1101\n")
		os.Exit(1)
	}

	binaryString := os.Args[1]
	result, err := modthree.ModThree(binaryString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The result of %s mod 3 is: %d\n", binaryString, result)
}
