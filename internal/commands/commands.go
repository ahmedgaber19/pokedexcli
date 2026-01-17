package commands

import "github.com/ahmedgaber19/pokedexcli/internal/pokedex"

// Command represents a CLI command with its metadata and callback
type Command struct {
	Name        string
	Description string
	Callback    func(*pokedex.Config, []string) error
}

// GetCommands returns a map of all available commands
func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays available commands",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays location areas",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous location areas",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a specific location area",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a specific Pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught Pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all caught Pokemon",
			Callback:    CommandPokedex,
		},
	}
}
