package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	// fmt.Printf("%v\n", pokemon)
	pokemonData, exists := cfg.pokedex[pokemon]
	if exists {
		fmt.Printf("Name: %v\n", pokemonData.Name)
		fmt.Printf("Height: %v\n", pokemonData.Height)
		fmt.Printf("Weight: %v\n", pokemonData.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemonData.Stats {
			fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types: \n")
		for _, x := range pokemonData.Types {
			fmt.Printf("\t- %v\n", x.Type.Name)
		}
	} else {
		fmt.Printf("you have not caught that pokemon\n")
	}

	return nil
}
