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

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// projectService, err := service.NewProjectService()
	// projects, err := projectService.GetAllProjects(ctx)

	// if err != nil {
	// 	fmt.Printf("Fehler beim Abrufen der Projekte: %v\n", err)
	// 	return
	// }

	// // Projekte ausgeben
	// for _, project := range projects {
	// 	fmt.Printf("Projekt: %s (UUID: %s)\n", project.Name, project.ID)
	// }

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

		//projectService := service.NewProjectService(projectRepo)

	}
}
