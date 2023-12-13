package main

import (
	"fmt"
)

func exploreCallback(cfg *config, areaName string) error {
	pokiClient := cfg.pokeApiClient
	area, err := pokiClient.ExploreArea(areaName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found Pokemon:")
	for _, v := range area.PokemonEncounters {

		fmt.Printf(" - %v \n", v.Pokemon.Name)
	}
	return nil

}
