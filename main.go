package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Response struct {
	Count int `json:"count"`
	Config
	Locations []Location `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
	`)
	return nil
}

func commandMap(config *Config) error {
	if config.Next == "" {
		return nil
	}
	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}
	for _, loc := range response.Locations {
		fmt.Println(loc.Name)
	}
	config.Next = response.Config.Next
	config.Previous = response.Previous
	return nil
}
func commandMapb(config *Config) error {
	if config.Previous == "" {
		return nil
	}
	res, err := http.Get(config.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}
	fmt.Println()
	for _, loc := range response.Locations {
		fmt.Println(loc.Name)
	}
	config.Next = response.Config.Next
	config.Previous = response.Previous
	return nil
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Show help information",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show previous locations",
		callback:    commandMapb,
	},
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
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	config := Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		line := sc.Text()
		command, ok := supportedCommands[line]
		if ok {
			command.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}

	}
}
