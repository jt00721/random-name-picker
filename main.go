package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Read names from file logic
func readNamesFromFile(filename string) ([]string, error) {
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var names []string
	// Initialise scanner to read from file
	scanner := bufio.NewScanner(file)
	// Scan() will remain true till the end of input (file) is reached or an error occurs
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	return names, scanner.Err()
}

func main() {
	// Provide usage of program if no names or file is provided as arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  Provide names as arguments: go run main.go Alice Bob Charlie")
		fmt.Println("  Or provide a file with names: go run main.go -file names.txt")
		return
	}
	// Read names from file provided as argument
	var names []string
	if os.Args[1] == "-file" {
		// Validate file path is provided
		if len(os.Args) < 3 {
			fmt.Println("Please provide the file path after -file.")
			return
		}

		filepath := os.Args[2]
		var err error
		names, err = readNamesFromFile(filepath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
	} else {
		// Slice of names from the provided arguments
		names = os.Args[1:]
	}

	// Check if no names are provided
	if len(names) == 0 {
		fmt.Println("No names provided.")
		return
	}

	// Check if only one name is provided
	if len(names) == 1 {
		fmt.Println("Why would you use me for ONE name??... Hmmm... Maybe the name you want is", names[0])
		return
	}

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
