package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

const DefaultLocationAreaApiUrl = "https://pokeapi.co/api/v2/location-area/"

func FetchLocationAreas(apiUrl string, cache *pokecache.Cache) (la LocationAreas, err error) {
	if data, ok := cache.Get(apiUrl); ok {
		if err = json.Unmarshal(data, &la); err != nil {
			return LocationAreas{}, fmt.Errorf("error unmarshaling cache data: %w", err)
		}
		return
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error fetching location area: %w", err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf("request wasn't successful, response: %s", resBody)
	}

	err = json.Unmarshal(resBody, &la)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	cache.Add(apiUrl, resBody)
	return
}
