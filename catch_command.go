package main

import (
	"fmt"
	"math/rand"
)

func catchCallBack(cfg *config, name string) error {
	client := cfg.pokeApiClient
	pokeMonDatabase := cfg.pokeMonsData

	pokeMon, err := client.GetPokeman(name)
	if err != nil {
		return fmt.Errorf("failed to get %v information: %v", name, err)
	}
	//after succesfuly getting the pokemon, lets save it to database
	pokeMonDatabase.AddPokeMon(name, pokeMon)
	r := rand.Int()
	if r <= pokeMon.BaseExperience {
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
