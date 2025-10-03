package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = Config{
	Next:     "",
	Previous: "",
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config)
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
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
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
	}
}

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	return strings.Fields(cleaned)
}
