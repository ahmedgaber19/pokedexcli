package pokeapi

// LocationAreaResult represents the response from the location-area API endpoint
type LocationAreaResult struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ExploreAreaResult represents detailed information about a specific location area
type ExploreAreaResult struct {
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

// Pokemon represents a Pokemon with its stats and attributes
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStats int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
