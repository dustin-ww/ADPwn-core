package states

import (
	"ADPwn/database/project/model"
	"ADPwn/tools"
	db_context "context"
	"fmt"
	"time"
)

type MainMenuState struct {
	Project model.Project
}

func (s *MainMenuState) Execute(context *Context) {
	fmt.Println("\nWelcome to main menu for project: " + s.Project.Name)

	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()
	/* projectService, _ := service.NewProjectService()
	projectService.SaveSubnet(ctx, s.Project, "192.192.56.4/30") */

	context.SetState(nil)

	fmt.Println("\nADPwn - Menu for project: " + s.Project.Name + " (" + s.Project.UID + ")")
	fmt.Println("--- Configure Project --- ")
	fmt.Println("1. ")
	fmt.Println("1. Run nmap")
	fmt.Println("2. Exit")
	var choice int
	fmt.Print("Please choose options: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		tools.Execute()
	case 2:
		fmt.Println("Exit...")
		context.SetState(nil)
	default:
		fmt.Println("Invalid option.")
	}
}
