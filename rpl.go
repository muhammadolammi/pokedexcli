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
	callback    func(*config, string) error
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
			callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 location area in game",
			callback:    mapbCallback,
		},
		"explore": {
			name:        "explore",
			description: "call this function as explore<area> to get pokemon in input areas. eg explore pastoria-city-area",
			callback:    exploreCallback,
		},
		"catch": {
			name:        "catch",
			description: "call this function as catch<name> to get catch the pokeman . eg catch pikachu",
			callback:    catchCallBack,
		},
		"inspect": {
			name:        "inspect",
			description: "call this function as inspect<name> to know if you have catched the pokeman . eg inspect pikachu",
			callback:    inspectCallback,
		},
		"pokedex": {
			name:        "pokedex",
			description: "call this function as to get all catched pokemans",
			callback:    pokedexCallback,
		},
	}

}

func Rpl(cfg *config) {
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
		area := ""
		if len(cleaned) > 1 {
			area = cleaned[1]
		}

		commands := getCommands()
		if _, ok := commands[command]; !ok {
			fmt.Println("Invalid Commad")
			continue
		}
		err := commands[command].callback(cfg, area)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanText(s string) []string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return words

}
