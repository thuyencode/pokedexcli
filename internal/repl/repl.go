package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thuyencode/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name, description string
	callback          func(*cliConfig) error
}

type cliConfig struct {
	commands     map[string]cliCommand
	locationArea *pokeapi.LocationArea
}

func Repl() {
	config := cliConfig{commands: map[string]cliCommand{
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
	}}

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

		command, ok := config.commands[args[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&config)
		if err != nil {
			fmt.Printf("Error executing %q command: %s\n", command.name, err.Error())
		}
	}
}
