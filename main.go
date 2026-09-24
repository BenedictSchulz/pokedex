package main

import (
	"time"

	"github.com/BenedictSchulz/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
