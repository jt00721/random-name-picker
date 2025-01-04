package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Check if no names provided as arguments
	if len(os.Args) < 2 {
		fmt.Println("No names provided")
		return
	}
	// Check if more than 2 names are provided as arguments
	if len(os.Args) < 3 {
		fmt.Println("Please provide more than one name as arguments.")
		return
	}

	// Declare a slice of names from the provided arguments
	names := os.Args[1:]

	// Print how many names were provided
	fmt.Println("Number of provided names:", len(names))

	// Initialise the default source to a deterministic state
	// If Seed() is not called the same pseudo-random sequence is generated
	// i.e. The same option (name) will be output
	rand.Seed(time.Now().UnixNano())

	// Randomly select a number from 0 to the length of names - 1
	// Use the randomly selected number as the index to use for the names slice
	selectedName := names[rand.Intn(len(names))]

	// Print the randomly selected namee
	fmt.Println("Randomly selected name:", selectedName)
}
