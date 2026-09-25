package repl

import (
	"errors"
	"fmt"
	"os"

	"github.com/thuyencode/pokedexcli/internal/pokeapi"
)

var ErrNoCommandRegistered = errors.New("no commands registered")

func commandExit(_ *cliConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(c *cliConfig) error {
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

func commandMap(c *cliConfig) error {
	fmt.Println("Fetching data...")
	fmt.Println()

	var apiUrl string

	if c.locationArea == nil || c.locationArea.Next == nil {
		apiUrl = pokeapi.DefaultLocationAreaApiUrl
	} else {
		apiUrl = *c.locationArea.Next
	}

	data, err := pokeapi.FetchLocationArea(apiUrl)
	if err != nil {
		return err
	}

	c.locationArea = &data

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(c *cliConfig) error {
	fmt.Println("Fetching data...")
	fmt.Println()

	var apiUrl string

	if c.locationArea == nil || c.locationArea.Previous == nil {
		apiUrl = pokeapi.DefaultLocationAreaApiUrl
	} else {
		apiUrl = *c.locationArea.Previous
	}

	data, err := pokeapi.FetchLocationArea(apiUrl)
	if err != nil {
		return err
	}

	c.locationArea = &data

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}
