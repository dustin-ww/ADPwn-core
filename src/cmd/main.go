package main

import (
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

	if len(progArgs) > 1 {
		fmt.Println("Unrecognized program options. Please type --help for valid arguments")
		os.Exit(1)
	}

	switch args := progArgs[1]; args {
	case "--version":
		fmt.Println("Version 0.0.1 Alpha")
	case "--help":
		fmt.Println("Type ADPwn start to enter program")
	case "start":
		startApp()
	default:
		fmt.Println("Unrecognized program options!")
		os.Exit(1)
	}
}

func startApp() {
	var lastInput string

	fmt.Println(asciiArt)

	for {
		fmt.Println("Type Option")
		fmt.Scan(&lastInput)
	}
}
