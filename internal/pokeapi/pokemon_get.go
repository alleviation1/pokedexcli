package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonDetails(pokemonName string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// attempt to get cached data
	if res, ok := c.cache.Get(url); ok {
		pokemon := PokemonDetails{}
		err := json.Unmarshal(res, &pokemon)
		if err != nil {
			return PokemonDetails{}, fmt.Errorf("Error unmarshaling cached pokemon data: %v", err)
		}

		return pokemon, nil
	}

	// default http call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Error reading pokemon details data: %v", err)
	}

	pokemon := PokemonDetails{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("Error unmarshaling data into pokemon details: %v", err)
	}

	c.cache.Add(url, data)
	return pokemon, nil
}

