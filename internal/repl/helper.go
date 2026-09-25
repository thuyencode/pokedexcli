package repl

import "strings"

func cleanInput(text string) []string {
	return strings.Fields(text)
}
