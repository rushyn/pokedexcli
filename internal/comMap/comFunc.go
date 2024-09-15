package comMap

import (
	"errors"
	"fmt"
	"os"

	"github.com/rushyn/pokedexcli/internal/PokeAPI"
)

func commandHelp(str string) error {
	if str != ""{
		fmt.Println("---------------------------------------------------")
		fmt.Println("help can't take arumets normal help return below")
		fmt.Println("---------------------------------------------------")
	}

	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

func commandExit(str string) error {
	if str == "" {
		os.Exit(0)
		return nil
	}else{
		return errors.New("exit command can't take an argument try entering >exit< no space and hit enter")
	}
}

func commandMap(str string) error {
	if str == ""{
		PokeAPI.GetLocations("next")
		return nil
	}else{
		return errors.New("map command can't take an argument")
	}
}

func commandMapb(str string) error {
	if str == ""{
		PokeAPI.GetLocations("back")
		return nil
	}else{
		return errors.New("mapb command can't take an argument")
	}
}

func commandExplore (str string) error {
	if str != ""{
		PokeAPI.GetPokemon(str)
		return nil
	}else{
		return errors.New("explore needs arear to explore")
	}
}