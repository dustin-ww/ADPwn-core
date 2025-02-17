package main

import (
	"ADPwn/cmd/internal/states"
	"ADPwn/cmd/internal/states/common"
	"fmt"
	"os"

	"github.com/rivo/tview"
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
	app := tview.NewApplication()

	context := common.Context{}
	context.SetState(&states.StartState{App: app})

	for context.CurrentState != nil {
		context.Execute()
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
