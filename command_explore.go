package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}

	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocationDetails(locationName)
	if err != nil {
		return fmt.Errorf("Error getting pokemon from location-area")
	}

	fmt.Println(locationName)
	fmt.Println("Found Pokemon:")
	for i := range(location.PokemonEncounters) {
		fmt.Printf("- %s\n", location.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}