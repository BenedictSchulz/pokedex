package main

import (
	"fmt"
	"math/rand"
)

func catchChance(level int) float64 {
	const (
		base      = 0.90
		perLevel  = 0.04
		minChance = 0.05
	)

	chance := base - perLevel*float64(level)
	if chance < minChance {
		chance = minChance
	}

	return chance
}

func tryCatch(level int) bool {
	return rand.Float64() < catchChance(level)
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catched := tryCatch(pokemon.BaseExperience)
	if catched == false {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
