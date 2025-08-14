package main

import (
	"time"

	"github.com/retrophozeac/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
