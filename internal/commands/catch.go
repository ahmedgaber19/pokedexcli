package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// CommandCatch attempts to catch a Pokemon
func CommandCatch(c *pokedex.Config, args []string) error {
	if len(args) <= 0 {
		return errors.New("please provide a Pokemon name to catch")
	}

	pokeName := args[0]
	_, ok := c.Pokemons[pokeName]
	if ok {
		fmt.Printf("%v was already caught!\n", pokeName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokeName)
	pokeURL := fmt.Sprintf(c.PokemonURL+"%v", pokeName)
	pokeRes, err := c.APIClient.GetPokemon(pokeURL)
	if err != nil {
		return err
	}

	prob := rand.Float32()
	if prob >= 0.5 {
		c.Pokemons[pokeName] = *pokeRes
		fmt.Printf("%v was caught!\n", pokeName)
	} else {
		fmt.Printf("%v escaped!\n", pokeName)
	}
	return nil
}
