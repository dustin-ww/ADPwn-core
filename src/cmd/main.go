package main

import (
	"ADPwn/cmd/states"
	"fmt"
	"os"
)

const asciiArt string = `
 /$$$$$$  /$$$$$$$  /$$$$$$$                         
 /$$__  $$| $$__  $$| $$__  $$                        
| $$  \ $$| $$  \ $$| $$  \ $$ /$$  /$$  /$$ /$$$$$$$ 
| $$$$$$$$| $$  | $$| $$$$$$$/| $$ | $$ | $$| $$__  $$
| $$__  $$| $$  | $$| $$____/ | $$ | $$ | $$| $$  \ $$
| $$  | $$| $$  | $$| $$      | $$ | $$ | $$| $$  | $$
| $$  | $$| $$$$$$$/| $$      |  $$$$$/$$$$/| $$  | $$
|__/  |__/|_______/ |__/       \_____/\___/ |__/  |__/
                                                      
                                                      
                                                    
`

func main() {

	progArgs := os.Args

	if len(progArgs) >= 2 {
		handleAdditionalProgramArgs(progArgs)
	}
	startApp()
}

func handleAdditionalProgramArgs(additionalArgs []string) {
	switch args := additionalArgs[1]; args {
	case "--version":
		fmt.Println("Version 0.0.1 Alpha")
	case "--help":
		fmt.Println("Type ADPwn start to enter program")
	case "start":
		startApp()
	default:
		fmt.Println("Unrecognized program options. Please type --help for valid arguments")
	}
	os.Exit(1)
}

func startApp() {
	fmt.Println(asciiArt)

	context := &states.Context{}
	context.SetState(&states.StartMenuState{})

	for context.CurrentState != nil {
		context.Execute()
	}
}
