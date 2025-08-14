package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon_name string) (Pokemon, error) {
	url := fmt.Sprintf("%v/pokemon/%v", baseURL, pokemon_name)
	// url :=  baseURL + "/location-area"
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		// fmt.Printf("using cache")
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)
	return pokemon, nil
}
