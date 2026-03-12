package main

import(
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must enter the pokemon's name")
	}

	pokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemon_details, err := cfg.pokeapiClient.GetPokemonDetails(pokemon)
	if err != nil {
		return fmt.Errorf("Error getting pokemon details in catch command: %v", err)
	}

	roll := float64(rand.Intn(99) + 1)
	catch_chance := 100.00 - (float64(pokemon_details.BaseExperience) * 0.25)
	fmt.Printf("roll: %v\n catch chance: %v\n", roll, catch_chance)
	if roll + catch_chance > 100.00 {
		fmt.Printf("%v was caught!\n", pokemon_details.Name)
		cfg.pokedex[pokemon_details.Name] = pokemon_details
	} else {
		fmt.Printf("%v escaped!\n", pokemon_details.Name)
	}

	return nil
}