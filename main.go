package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type LocationAreasResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
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
		"map": {
			name:        "map",
			description: "Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func main() {
	cfg := &config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

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
			err := cmd.callback(cfg)
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandMap(cfg *config) error {
	if cfg.Next == "" {
		fmt.Println("No more locations available")
		return nil
	}

	resp, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var locationAreas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	// Update config with new pagination URLs
	if locationAreas.Next != nil {
		cfg.Next = *locationAreas.Next
	} else {
		cfg.Next = ""
	}
	if locationAreas.Previous != nil {
		cfg.Previous = *locationAreas.Previous
	} else {
		cfg.Previous = ""
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var locationAreas LocationAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	// Update config with new pagination URLs
	if locationAreas.Next != nil {
		cfg.Next = *locationAreas.Next
	} else {
		cfg.Next = ""
	}
	if locationAreas.Previous != nil {
		cfg.Previous = *locationAreas.Previous
	} else {
		cfg.Previous = ""
	}

	return nil
}
