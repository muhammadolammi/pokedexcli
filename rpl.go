package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"map": {
			name:        "map",
			description: "Get next 20 location area in game",
			callback:    func() error { return nil },
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 location area in game",
			callback: func() error {
				return nil
			},
		}}
}

func Rpl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {

			continue
		}
		cleaned := cleanText(text)
		command := cleaned[0]
		commands := getCommands()
		if _, ok := commands[command]; !ok {
			fmt.Println("Invalid Commad")
			continue
		}
		commands[command].callback()

	}
}

func cleanText(s string) []string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return words

}
