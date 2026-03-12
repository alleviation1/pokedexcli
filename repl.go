package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"

	"github.com/alleviation1/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {


	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"mapf": {
			name:		 "mapf",
			description: "Display 20 Pokemon map locations",
			callback:	 commandMapf,
		},
		"mapb": {
			name:		 "mapb",
			description: "Display previous 20 Pokemon map locations",
			callback:	 commandMapb,
		},
		"explore": {
			name:		  "explore <location_name>",
			description:  "Display all Pokemon in a location",
			callback:	  commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon_name>",
			description: "attempt to catch a pokemon",
			callback:	 commandCatch,
		},
	}
}

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		fmt.Printf("Your command was: %s\n", commandName)

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...); if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}