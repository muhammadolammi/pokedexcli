package main

import (
	"time"

	"github.com/muhammadolammi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	next          *string
	previous      *string
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
	}
	Rpl(&cfg)

}
