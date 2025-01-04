package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if more than 2 names are provided as arguments
	if len(os.Args) < 3 {
		fmt.Println("Please provide more than one name as arguments.")
		return
	}

	// Declare a slice of names from the provided arguments
	names := os.Args[1:]

	// Print the provided names
	fmt.Println("Names provided:", names)
}
