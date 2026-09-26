package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/thuyencode/pokedexcli/internal/pokecache"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var Client = http.Client{Timeout: 10 * time.Second}

const DefaultLocationAreaApiUrl = "https://pokeapi.co/api/v2/location-area/"

func FetchLocationAreas(apiUrl string, cache *pokecache.Cache) (LocationAreas, error) {
	var las LocationAreas

	if data, ok := cache.Get(apiUrl); ok {
		if err := json.Unmarshal(data, &las); err != nil {
			return LocationAreas{}, fmt.Errorf("error unmarshaling cache data: %w", err)
		}
		return las, nil
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error fetching location areas: %w", err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf("request wasn't successful, response: %q", resBody)
	}

	err = json.Unmarshal(resBody, &las)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	cache.Add(apiUrl, resBody)

	return las, nil
}

func FetchLocationArea(apiUrl string, cache *pokecache.Cache) (LocationArea, error) {
	var la LocationArea

	if data, ok := cache.Get(apiUrl); ok {
		if err := json.Unmarshal(data, &la); err != nil {
			return LocationArea{}, fmt.Errorf("error unmarshaling cache data: %w", err)
		}
		return la, nil
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error fetching location area: %w", err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading response body: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("request wasn't successful, response: %q", resBody)
	}

	err = json.Unmarshal(resBody, &la)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	cache.Add(apiUrl, resBody)
	return la, nil
}
