package comMap

import (
	"fmt"
	"os"

	"github.com/rushyn/pokedexcli/internal/PokeAPI"
)

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	PokeAPI.GetLocations("next")
	
	return nil
}

func commandMapb() error {
	PokeAPI.GetLocations("back")
	return nil
}