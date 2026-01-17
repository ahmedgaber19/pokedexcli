package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// CommandMap displays the next page of location areas
func CommandMap(c *pokedex.Config, _ []string) error {
	if c.Next == "" {
		return errors.New("no more location areas available")
	}

	cachedVal, found := c.Cache.Get(c.Next)
	var locationResult *struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	if !found {
		result, err := c.APIClient.GetLocationAreas(c.Next)
		if err != nil {
			return err
		}

		locationResult = &struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Results  []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"results"`
		}{
			Next:     result.Next,
			Previous: result.Previous,
			Results:  result.Results,
		}

		byteVal, err := json.Marshal(locationResult)
		if err != nil {
			return err
		}
		c.Cache.Add(c.Next, byteVal)
	} else {
		locationResult = &struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Results  []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"results"`
		}{}
		if err := json.Unmarshal(cachedVal, locationResult); err != nil {
			return err
		}
	}

	c.Next = locationResult.Next
	c.Previous = locationResult.Previous

	for _, location := range locationResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

// CommandMapb displays the previous page of location areas
func CommandMapb(c *pokedex.Config, _ []string) error {
	if c.Previous == "" {
		return errors.New("no previous location areas available")
	}

	cachedVal, found := c.Cache.Get(c.Previous)
	var locationResult *struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	if !found {
		result, err := c.APIClient.GetLocationAreas(c.Previous)
		if err != nil {
			return err
		}

		locationResult = &struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Results  []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"results"`
		}{
			Next:     result.Next,
			Previous: result.Previous,
			Results:  result.Results,
		}

		byteVal, err := json.Marshal(locationResult)
		if err != nil {
			return err
		}
		c.Cache.Add(c.Previous, byteVal)
	} else {
		locationResult = &struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Results  []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"results"`
		}{}
		if err := json.Unmarshal(cachedVal, locationResult); err != nil {
			return err
		}
	}

	c.Next = locationResult.Next
	c.Previous = locationResult.Previous

	for _, location := range locationResult.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
