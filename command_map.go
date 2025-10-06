package main

import (
	"fmt"

	pokeapi "github.com/kyleludlow/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Next) > 0 {
		url = config.Next
	}
	res, err := pokeapi.GetData[pokeapi.PokeResponseWrapper[pokeapi.LocationArea]](url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	if res.Previous != nil && len(*res.Previous) > 0 {
		config.Previous = *res.Previous
	}
	if res.Next != nil && len(*res.Next) > 0 {
		config.Next = *res.Next
	}
	for _, location := range res.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapb(config *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Previous) > 0 {
		url = config.Previous
	}

	res, err := pokeapi.GetData[pokeapi.PokeResponseWrapper[pokeapi.LocationArea]](url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	if res.Previous != nil && len(*res.Previous) > 0 {
		config.Previous = *res.Previous
	}
	if res.Next != nil && len(*res.Next) > 0 {
		config.Next = *res.Next
	}
	for _, location := range res.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
