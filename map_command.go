package main

import (
	"fmt"
)

func mapCallback(cfg *config) error {
	pokiClient := cfg.pokeApiClient
	res, err := pokiClient.MakeRequest(cfg.next)
	if err != nil {
		fmt.Printf("there is an error : %v\n", err)
	}
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	cfg.next = res.Next

	return nil

}
