package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rushyn/pokedexcli/internal/comMap"
)



func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		valCli, exists := comMap.ComMap[scanner.Text()]
		if exists {
			err := valCli.Callback()
			if err != nil {
				fmt.Printf("Some is wrong err is ::: %v", err)
			}
		} else {
			println("\n!!! Invalid Command !!! try >help<\n")
		}
		fmt.Print("Pokedex > ")
	}
}
