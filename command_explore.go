package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemons(location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\nFound Pokemon:\n", location)
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
