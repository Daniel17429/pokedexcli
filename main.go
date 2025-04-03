package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

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

		if cmd, ok := commands[command]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
