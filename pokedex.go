package main

import (
	"errors"
	"fmt"
)

func pokedexCallback(cfg *config, name string) error {
	pokes := cfg.pokeMonsData
	if pokes == nil {
		return errors.New("no Pokemans yet, try catching new ones")
	}
	for _, poke := range pokes {
		fmt.Printf(" - %v \n", poke.Name)
	}
	return nil
}
