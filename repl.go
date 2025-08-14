package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/retrophozeac/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		if (commandName == "explore" || commandName == "catch" || commandName == "inspect") && len(words) < 2 {
			fmt.Printf("need more input")
			continue
		}
		input := ""
		if commandName == "explore" || commandName == "catch" || commandName == "inspect" {
			input = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknow Command\n")
			continue
		}

	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, input string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays 20 location areas in Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list pokemons in the location-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects caught pokemon",
			callback:    commandInspect,
		},
	}
}

type config struct {
	pokedex          map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}
