package repl

import (
	"errors"
	"fmt"
	"os"
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
