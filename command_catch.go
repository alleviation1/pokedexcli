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

	roll := rand.Intn(100)
	catch_chance := 50.0 / float32(pokemon_details.BaseExperience)
	fmt.Println(catch_chance * 100.0)
	if float32(roll) > catch_chance {
		fmt.Println("Caught!")
	} else {
		fmt.Println("Not caught!")
	}

	return nil
}