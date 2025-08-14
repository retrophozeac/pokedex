package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, pokemon string) error {
	pokemons, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	base_exp := pokemons.BaseExperience
	caught := catchPokemon(base_exp, 608)
	if caught {
		fmt.Printf("%v was caught!\n", pokemon)
		cfg.pokedex[pokemon] = pokemons
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}
	return nil
}
func catchPokemon(baseExp int, maxBaseExp int) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	probability := float64(maxBaseExp-baseExp) / float64(maxBaseExp)
	return r.Float64() < probability
}
