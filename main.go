package main

import (
	"time"

	"github.com/maker2413/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapiClient: pokeClient,
	}

	startRepl(&cfg)
}
