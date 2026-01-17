package commands

import (
	"errors"
	"fmt"

	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// CommandInspect displays detailed information about a caught Pokemon
func CommandInspect(c *pokedex.Config, args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a Pokemon name to inspect")
	}

	pokeName := args[0]
	poke, ok := c.Pokemons[pokeName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Println("Stats:")
	for _, s := range poke.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStats)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}

// CommandPokedex lists all caught Pokemon
func CommandPokedex(c *pokedex.Config, _ []string) error {
	if len(c.Pokemons) < 1 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, poke := range c.Pokemons {
		fmt.Printf(" - %s\n", poke.Name)
	}
	return nil
}
