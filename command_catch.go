package main

import (
	"fmt"

	"math/rand"

	pokeapi "github.com/kyleludlow/pokedexcli/internal/pokeapi"
)

func commandCatch(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + param

	res, err := pokeapi.GetData[pokeapi.PokemonRes](url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	chanceMultiplier := res.BaseExperience
	catchRoll := rand.Intn(chanceMultiplier)
	isCaught := catchRoll > (chanceMultiplier / 2)

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	if isCaught {
		fmt.Printf("%s was caught!\n", res.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		config.Pokedex[res.Name] = res
	} else {
		fmt.Printf("%s was NOT caught!\n", res.Name)
	}

	return nil
}
