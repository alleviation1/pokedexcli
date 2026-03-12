package main

import(
	"time"

	"github.com/alleviation1/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: 	make(map[string]pokeapi.PokemonDetails),
	}

	startRepl(cfg)
}