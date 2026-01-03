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

	"github.com/ahmedgaber19/pokedexcli/internal/pokecache"
)

type Config struct {
	Next       string `json:"next"`
	Previous   string `json:"previous"`
	cache      *pokecache.Cache
	exploreUrl string
}

type LocationAreaResult struct {
	Config
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ExploreAreaResult struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

var httpCli = http.Client{
	Timeout: time.Second * 30,
}

func commandExist(_ *Config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommandsMap() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *Config, _ []string) error {
	if c.Next == "" {
		return errors.New(" No more location areas available.")
	}
	var LocationAreaResult LocationAreaResult
	cachedVal, found := c.cache.Get(c.Next)
	if !found {
		res, err := httpCli.Get(c.Next)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&LocationAreaResult)
		if err != nil {
			return err
		}
		byteVal, err := json.Marshal(LocationAreaResult)
		if err != nil {
			return err
		}
		c.cache.Add(c.Next, byteVal)

	} else {
		json.Unmarshal(cachedVal, &LocationAreaResult)
	}
	c.Next = LocationAreaResult.Next
	c.Previous = LocationAreaResult.Previous
	for _, location := range LocationAreaResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil

}

func commandMapb(c *Config, _ []string) error {
	if c.Previous == "" {
		return errors.New("No previous location areas available.")
	}
	var LocationAreaResult LocationAreaResult
	cachedVal, found := c.cache.Get(c.Previous)

	if !found {
		res, err := httpCli.Get(c.Previous)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&LocationAreaResult)
		if err != nil {
			return err
		}
		byteVal, err := json.Marshal(LocationAreaResult)
		if err != nil {
			return err
		}
		c.cache.Add(c.Previous, byteVal)
	} else {
		json.Unmarshal(cachedVal, &LocationAreaResult)
	}
	c.Next = LocationAreaResult.Next
	c.Previous = LocationAreaResult.Previous
	for _, location := range LocationAreaResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func connamdExplore(c *Config, args []string) error {
	if len(args) <= 0 {
		return errors.New("Please provide a location area name to explore.")
	}
	var exploreRes ExploreAreaResult
	cityName := args[0]
	cachedVal, found := c.cache.Get(cityName)
	if !found {
		url := fmt.Sprintf(c.exploreUrl+"%v", cityName)
		res, err := httpCli.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&exploreRes)
		if err != nil {
			return err
		}
		byteVal, err := json.Marshal(exploreRes)
		if err == nil {
			c.cache.Add(cityName, byteVal)
		}

	} else {
		err := json.Unmarshal(cachedVal, &exploreRes)
		if err != nil {
			return err
		}
	}

	if len(exploreRes.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in %v.\n", cityName)
		return nil
	}
	fmt.Printf("Exploring %v...\n", cityName)
	fmt.Println("Found Pokemon:")
	for _, pokes := range exploreRes.PokemonEncounters {
		fmt.Printf("- %v\n", pokes.Pokemon.Name)
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
			callback:    commandHelp,
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
		"explore": {
			name:        "explore",
			description: "Explore a specific location area",
			callback:    connamdExplore,
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
	pokeCache := pokecache.NewCache(time.Second * 5)
	c := Config{
		Next:       "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous:   "",
		cache:      pokeCache,
		exploreUrl: "https://pokeapi.co/api/v2/location-area/",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerText := scanner.Text()
		words := cleanInput(scannerText)
		if len(words) == 0 {
			continue
		}
		firstWord := words[0]
		command, ok := commandMap[firstWord]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}
		err := command.callback(&c, words[1:])
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

	}
}
