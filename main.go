package main

import (
	// "fmt"
	"time"

	"github.com/muhammadolammi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeMonsData  pokeapi.MyPokeMons
	pokeApiClient pokeapi.Client
	next          *string
	previous      *string
}

func main() {
	cfg := config{
		pokeMonsData:  pokeapi.NewPokeMonsData(),
		pokeApiClient: pokeapi.NewClient(time.Hour),
	}
	Rpl(&cfg)

	// fmt.Println(area)

}
