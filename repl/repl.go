package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ahmedgaber19/pokedexcli/internal/commands"
	"github.com/ahmedgaber19/pokedexcli/internal/pokedex"
)

// Start begins the REPL (Read-Eval-Print Loop) for the Pokedex CLI
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := commands.GetCommands()
	config := pokedex.NewConfig()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerText := scanner.Text()
		words := CleanInput(scannerText)

		if len(words) == 0 {
			continue
		}

		firstWord := words[0]
		command, ok := commandMap[firstWord]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}

		err := command.Callback(config, words[1:])
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

// CleanInput converts input text to lowercase words
func CleanInput(text string) []string {
	words := strings.Fields(text)
	res := make([]string, len(words))
	for i, word := range words {
		res[i] = strings.ToLower(word)
	}
	return res
}
