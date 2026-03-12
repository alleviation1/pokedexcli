package main

import (
	"fmt"
	"errors"
)

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("Error getting location areas in map forward: %w", err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range(locations.Results) {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("This is the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("Error getting location areas in map back: %w", err)
	}
	
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range(locations.Results) {
		fmt.Println(location.Name)
	}

	return nil
}