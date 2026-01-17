package commands

import (
	"fmt"
	"os"

	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// CommandExit exits the Pokedex application
func CommandExit(_ *pokedex.Config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// CommandHelp displays all available commands
func CommandHelp(_ *pokedex.Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
