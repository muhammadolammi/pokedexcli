package main

import (
	"errors"
	"fmt"
)

func mapbCallback(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("this is the first page, nothing to print")
	}
	pokiClient := cfg.pokeApiClient
	res, err := pokiClient.MakeRequest(cfg.previous)
	if err != nil {
		return err
	}
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	cfg.next = res.Next
	cfg.previous = res.Previous
	return nil

}
