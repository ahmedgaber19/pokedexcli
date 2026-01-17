package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// CommandExplore explores a specific location area and shows available Pokemon
func CommandExplore(c *pokedex.Config, args []string) error {
	if len(args) <= 0 {
		return errors.New("please provide a location area name to explore")
	}

	cityName := args[0]
	cachedVal, found := c.Cache.Get(cityName)

	var exploreRes *struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		GameIndex int    `json:"game_index"`
		Location  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}

	if !found {
		url := fmt.Sprintf(c.ExploreURL+"%v", cityName)
		result, err := c.APIClient.GetLocationArea(url)
		if err != nil {
			return err
		}

		exploreRes = &struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			GameIndex int    `json:"game_index"`
			Location  struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"location"`
			PokemonEncounters []struct {
				Pokemon struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"pokemon"`
			} `json:"pokemon_encounters"`
		}{
			ID:                result.ID,
			Name:              result.Name,
			GameIndex:         result.GameIndex,
			Location:          result.Location,
			PokemonEncounters: result.PokemonEncounters,
		}

		byteVal, err := json.Marshal(exploreRes)
		if err == nil {
			c.Cache.Add(cityName, byteVal)
		}
	} else {
		exploreRes = &struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			GameIndex int    `json:"game_index"`
			Location  struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"location"`
			PokemonEncounters []struct {
				Pokemon struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"pokemon"`
			} `json:"pokemon_encounters"`
		}{}
		if err := json.Unmarshal(cachedVal, exploreRes); err != nil {
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
