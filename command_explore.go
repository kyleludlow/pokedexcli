package main

import (
	"fmt"

	pokeapi "github.com/kyleludlow/pokedexcli/internal/pokeapi"
)

func commandExplore(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + param

	if len(config.Next) > 0 {
		url = config.Next
	}
	res, err := pokeapi.GetData[pokeapi.LocationArea](url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("%s\n", encounter.Pokemon.Name)
	}

	return nil
}
