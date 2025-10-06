package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/kyleludlow/pokedexcli/internal/pokecache"
)

func get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers and make request
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type PokeResponseWrapper[T any] struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T
}

func GetData[T any](url string) (T, error) {
	var ch = pokecache.CACHE

	dataStruct := new(T)

	// use cache if exists
	data, found := ch.Get(url)
	if found {
		jsonErr := json.Unmarshal(data, dataStruct)
		if jsonErr != nil {
			return *dataStruct, jsonErr
		}
		return *dataStruct, nil
	}

	body, err := get(url)
	if err != nil {
		return *dataStruct, err
	}

	// add bytes to cache
	ch.Add(url, body)

	jsonErr := json.Unmarshal(body, dataStruct)
	if jsonErr != nil {
		return *dataStruct, jsonErr
	}

	return *dataStruct, nil
}
