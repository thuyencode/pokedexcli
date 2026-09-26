package repl

import (
	"errors"
	"fmt"
	"os"

	"github.com/thuyencode/pokedexcli/internal/pokeapi"
)

var ErrNoCommandRegistered = errors.New("no commands registered")

func commandExit(_ *cliConfig, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(c *cliConfig, _ ...string) error {
	if len(c.commands) == 0 {
		return ErrNoCommandRegistered
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range c.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(c *cliConfig, _ ...string) error {
	fmt.Println("Fetching data...")
	fmt.Println()

	var apiUrl string

	if c.locationAreas == nil || c.locationAreas.Next == nil {
		apiUrl = pokeapi.DefaultLocationAreaApiUrl
	} else {
		apiUrl = *c.locationAreas.Next
	}

	data, err := pokeapi.FetchLocationAreas(apiUrl, c.cache)
	if err != nil {
		return err
	}

	c.locationAreas = &data

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(c *cliConfig, _ ...string) error {
	fmt.Println("Fetching data...")
	fmt.Println()

	var apiUrl string

	if c.locationAreas == nil || c.locationAreas.Previous == nil {
		apiUrl = pokeapi.DefaultLocationAreaApiUrl
	} else {
		apiUrl = *c.locationAreas.Previous
	}

	data, err := pokeapi.FetchLocationAreas(apiUrl, c.cache)
	if err != nil {
		return err
	}

	c.locationAreas = &data

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(c *cliConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("not enough argument(s)")
	}

	city := args[0]
	apiUrl := pokeapi.DefaultLocationAreaApiUrl + city

	fmt.Printf("Exploring %s...\n", city)
	fmt.Println()

	data, err := pokeapi.FetchLocationArea(apiUrl, c.cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokémon:")
	for _, e := range data.PokemonEncounters {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}

	return nil
}
