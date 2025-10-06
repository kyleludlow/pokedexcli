package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kyleludlow/pokedexcli/internal/pokeapi"
)

var config = Config{
	Next:     "",
	Previous: "",
	Pokedex:  make(map[string]pokeapi.PokemonRes),
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		param := ""
		if len(words) > 1 {
			param = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config, param)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type Config struct {
	Next     string
	Previous string
	Pokedex  map[string]pokeapi.PokemonRes
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, param string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "lists areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists previous areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	return strings.Fields(cleaned)
}
