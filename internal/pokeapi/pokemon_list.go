package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (PokemonEncounters, error) {
	url := fmt.Sprintf("%v/location-area/%v", baseURL, location)
	// url :=  baseURL + "/location-area"
	if val, ok := c.cache.Get(url); ok {
		pokemons := PokemonEncounters{}
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return PokemonEncounters{}, err
		}
		// fmt.Printf("using cache")
		return pokemons, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}
	pokemons := PokemonEncounters{}
	err = json.Unmarshal(dat, &pokemons)
	if err != nil {
		return PokemonEncounters{}, err
	}
	c.cache.Add(url, dat)
	return pokemons, nil
}
