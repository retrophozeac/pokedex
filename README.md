# Pokedex CLI

A command-line interface (CLI) application for interacting with the Pokedex. This tool allows you to explore the world of Pokemon, catch them, and view your collection, all from your terminal.

## Features

*   **Explore Pokemon World**: Discover new areas and the Pokemon that inhabit them.
*   **Catch Pokemon**: Encounter and catch Pokemon in the wild.
*   **Manage Your Pokedex**: View the Pokemon you have caught and inspect their details.
*   **Interactive REPL**: An interactive Read-Eval-Print Loop (REPL) to enter commands.

## Available Commands

*   `map`: Displays the next 20 locations in the Pokemon world.
*   `mapb`: Displays the previous 20 locations.
*   `explore <location_name>`: Shows a list of Pokemon in a given location.
*   `catch <pokemon_name>`: Attempts to catch a Pokemon.
*   `inspect <pokemon_name>`: View details about a caught Pokemon (if it's in your Pokedex).
*   `pokedex`: Lists all the Pokemon you have caught.
*   `help`: Displays a help message with all available commands.
*   `exit`: Exits the Pokedex CLI.

## Usage

To run the application, you can build and run the executable:

```sh
go build
./pokedexcli
```

Alternatively, you can run it directly:

```sh
go run .
```

## Dependencies

This project uses Go modules to manage dependencies. The main dependency is the PokeAPI.
