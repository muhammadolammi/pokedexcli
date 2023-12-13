package main

import "fmt"

func helpCallback(cfg *config, area string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Help Menu!")

	fmt.Println("Here are your available commads")
	commads := getCommands()
	for _, cmd := range commads {
		fmt.Printf(" - %v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil

}
