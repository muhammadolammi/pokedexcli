package main

import (
	"errors"
	"fmt"
)

func inspectCallback(cfg *config, name string) error {
	pokeMonsDatabase := cfg.pokeMonsData
	v, ok := pokeMonsDatabase[name]
	if !ok {

		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", v.Name)
	fmt.Printf("Height: %v\n", v.Height)
	fmt.Printf("Weight: %v\n", v.Weight)
	fmt.Printf("Stats: \n")
	for _, s := range v.Stats {
		fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range v.Types {
		fmt.Printf(" -%v \n", t.Type.Name)
	}
	return nil
}
