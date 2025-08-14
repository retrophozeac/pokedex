package main

import "fmt"

func commandpokedex(cfg *config, input string) error {
	fmt.Printf("Your Pokedex:\n")
	for key, _ := range cfg.pokedex {
		fmt.Printf("\t-%v\n", key)
	}
	return nil
}
