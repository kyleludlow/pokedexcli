package main

import "fmt"

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Next) > 0 {
		url = config.Next
	}
	res, err := getData[LocationArea](url)
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

func commandMapb(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(config.Previous) > 0 {
		url = config.Previous
	}
	res, err := getData[LocationArea](url)
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
