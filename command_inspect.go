package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must enter the pokemon's name")
	}

	pokemon := args[0]

	pokemon_details, ok := cfg.pokedex[pokemon]
	if !ok {
		return errors.New("You have not caught that pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon_details.Name)
	fmt.Printf("Height: %d\n", pokemon_details.Height)
	fmt.Printf("Weight: %d\n", pokemon_details.Weight)


	fmt.Println("Stats:")
	for _, stat := range(pokemon_details.Stats) {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: \n")
	for _, pokemonType := range(pokemon_details.Types) {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
	return nil
}