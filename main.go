package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type LocationAreaResult struct {
	Config
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var httpCli = http.Client{
	Timeout: time.Second * 30,
}

func commandExist(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommandsMap() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *Config) error {
	if c.Next == "" {

		return errors.New("")
	}
	res, err := httpCli.Get(c.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var LocationAreaResult LocationAreaResult
	err = json.NewDecoder(res.Body).Decode(&LocationAreaResult)
	if err != nil {
		return nil
	}
	c.Next = LocationAreaResult.Next
	c.Previous = LocationAreaResult.Previous
	for _, location := range LocationAreaResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil

}

func commandMapb(c *Config) error {
	if c.Previous == "" {
		return errors.New("No previous location areas available.")
	}
	res, err := httpCli.Get(c.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var LocationAreaResult LocationAreaResult
	err = json.NewDecoder(res.Body).Decode(&LocationAreaResult)
	if err != nil {
		return nil
	}
	c.Next = LocationAreaResult.Next
	c.Previous = LocationAreaResult.Previous
	for _, location := range LocationAreaResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExist,
		},
		"help": {
			name:        "help",
			description: "Displays available commands",
			callback:    func(c *Config) error { return commandHelp(c) },
		},
		"map": {
			name:        "map",
			description: "Displays location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous location areas",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	res := make([]string, len(words))
	for i, word := range words {
		res[i] = strings.ToLower(word)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommandsMap()
	c := Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerText := scanner.Text()
		words := cleanInput(scannerText)
		firstWord := words[0]
		command, ok := commandMap[firstWord]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}
		err := command.callback(&c)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

	}
}
