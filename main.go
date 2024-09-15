package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rushyn/pokedexcli/internal/comMap"
)



func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		rawInput := scanner.Text()
		rawInput = strings.TrimSpace(rawInput)
		var input = strings.SplitN(rawInput, " ", 2)
		if len(input) == 1{
			input = append(input, "")
		}
		for i := range input{
			input[i] =strings.TrimSpace(input[i])
		}
	
		
		valCli, exists := comMap.ComMap[input[0]]

		if exists {
			err := valCli.Callback(input[1])
			if err != nil {
				fmt.Println("Some is wrong err is ::: ", err)
			}
		} else {
			println("\n!!! Invalid Command !!! try >help<\n")
		}
		fmt.Print("Pokedex > ")
	}
}
