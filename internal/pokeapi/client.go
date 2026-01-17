package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// Client handles HTTP requests to the PokeAPI
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new PokeAPI client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// GetLocationAreas fetches location areas from the given URL
func (c *Client) GetLocationAreas(url string) (*LocationAreaResult, error) {
	var result LocationAreaResult
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLocationArea fetches details about a specific location area
func (c *Client) GetLocationArea(url string) (*ExploreAreaResult, error) {
	var result ExploreAreaResult
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPokemon fetches details about a specific Pokemon
func (c *Client) GetPokemon(url string) (*Pokemon, error) {
	var result Pokemon
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
