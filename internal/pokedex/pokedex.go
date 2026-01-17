package pokedex

import (
	"time"

	"github.com/ahmedgaber19/pokedexcli/internal/pokeapi"
	"github.com/ahmedgaber19/pokedexcli/internal/pokecache"
)

// Config holds the state of the Pokedex application
type Config struct {
	Next       string
	Previous   string
	Cache      *pokecache.Cache
	ExploreURL string
	PokemonURL string
	Pokemons   map[string]pokeapi.Pokemon
	APIClient  *pokeapi.Client
}

// NewConfig creates a new Pokedex configuration
func NewConfig() *Config {
	return &Config{
		Next:       "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous:   "",
		Cache:      pokecache.NewCache(time.Second * 5),
		ExploreURL: "https://pokeapi.co/api/v2/location-area/",
		PokemonURL: "https://pokeapi.co/api/v2/pokemon/",
		Pokemons:   make(map[string]pokeapi.Pokemon),
		APIClient:  pokeapi.NewClient(),
	}
}
