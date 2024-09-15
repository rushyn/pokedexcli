package comMap



type cliCommand struct {
	name        string
	description string
	Callback    func(string) error
}


var ComMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		Callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Displays the names of next 20 location areas",
		Callback:   commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the names of previous 20 location areas",
		Callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Find the pokemoan",
		Callback:    commandExplore,
	},
	"catch": {
		name:        "catch ",
		description: "Catch the pokemoan",
		Callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect  ",
		description: "inspect the pokemoan",
		Callback:    commandInspect ,
	},
	
	
}

