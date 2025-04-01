package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Print prompt
		fmt.Print("Pokedex > ")

		// Wait for and read input
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		// Get and clean input
		input := scanner.Text()
		cleaned := cleanInput(input)

		// Handle empty input
		if len(cleaned) == 0 {
			continue
		}

		// Get first command
		command := cleaned[0]

		// Print command (for now)
		fmt.Printf("Your command was: %s\n", command)

		// Add exit condition
		if command == "exit" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
func cleanInput(text string) []string {
	// Trim leading and trailing whitespace
	trimmed := strings.TrimSpace(text)
	// Convert to lowercase
	lower := strings.ToLower(trimmed)
	// Split into words (splits on any whitespace)
	words := strings.Fields(lower)
	return words
}
