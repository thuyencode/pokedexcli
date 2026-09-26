package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/thuyencode/pokedexcli/internal/pokeapi"
	"github.com/thuyencode/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name, description string
	callback          func(c *cliConfig, args ...string) error
}

type cliConfig struct {
	commands      map[string]cliCommand
	locationAreas *pokeapi.LocationAreas
	cache         *pokecache.Cache
}

func Repl() {
	config := cliConfig{
		cache: pokecache.NewCache(10 * time.Second),
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Display a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Display next location areas",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Display previous location areas",
				callback:    commandMapBack,
			},
			"explore": {
				name:        "explore",
				description: "List all the Pokémon in a location area",
				callback:    commandExplore,
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				println("error reading user input: %w", err)
				return
			}

			return
		}

		input := strings.ToLower(scanner.Text())
		args := cleanInput(input)

		if len(args) == 0 {
			continue
		}

		commandName := args[0]
		rest := args[1:]
		command, ok := config.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&config, rest...)
		if err != nil {
			fmt.Printf("Error executing %q command: %s\n", command.name, err.Error())
		}
	}
}
