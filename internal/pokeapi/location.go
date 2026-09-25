package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const DefaultLocationAreaApiUrl = "https://pokeapi.co/api/v2/location-area/"

func FetchLocationArea(apiUrl string) (locationArea LocationArea, err error) {
	res, err := http.Get(apiUrl)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error fetching location area: %w", err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error deserializing response body: %w", err)
	}

	return
}
