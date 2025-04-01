package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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
