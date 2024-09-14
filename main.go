package main

import (
	"bufio"
	"fmt"
	"os"
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	PokeAPI.test()
	comMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		valCli, exists := comMap[scanner.Text()]
		if exists {
			err := valCli.callback()
			if err != nil {
				fmt.Printf("Some is wrong err is ::: %v", err)
			}
		} else {
			println("\n!!! Invalid Command !!! try >help<\n")
		}
		fmt.Print("Pokedex > ")
	}
}
