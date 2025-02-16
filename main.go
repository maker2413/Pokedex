package main

import (
	"time"

	"github.com/maker2413/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(&cfg)
}
