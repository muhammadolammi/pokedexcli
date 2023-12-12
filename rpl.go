package main

import "fmt"

func commandHelp() {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map:  Get next 20 location area in game
mapb: Get previous 20 location area in game
			`)

}
func commandExit() {
	return
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var commandsMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Get next 20 location area in game",
		callback:    func() {},
	},
	"mapb": {
		name:        "mapb",
		description: "Get previous 20 location area in game",
		callback: func() {

		},
	},
}

func Rpl() {
	for {
		fmt.Print("pokedex > ")
		var command string
		fmt.Scan((&command))
		switch command {
		case "help":
			commandsMap[command].callback()
			continue
		case "exit":
			commandsMap[command].callback()
			return
		case "map":
			commandsMap[command].callback()
		case "mapb":
			commandsMap[command].callback()
		default:
			fmt.Println("Wrong command, run help for help.")
			continue

		}
	}
}
