package states

import (
	"ADPwn/database/project/model"
	"fmt"
)

type MainMenuAddUserState struct {
	Project model.Project
}

func (s *MainMenuAddUserState) Execute(context *Context) {

	/* projectService, _ := service.NewProjectService()
	projectService.SaveSubnet(ctx, s.Project, "192.192.56.4/30") */

	context.SetState(nil)

	fmt.Println("\nADPwn - Menu for project: " + s.Project.Name + " (" + s.Project.UID + ")")
	fmt.Println("--- ADPwn add User --- ")
	fmt.Println("1. Add user with username")
	fmt.Println("2. Add user with ntlm hash")
	fmt.Println("3. Add user with kerberos ticket")
	fmt.Println("---------------------- ")
	fmt.Println("4. Back to main menu")
	var choice int
	fmt.Print("Please choose options: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println()
	case 2:
		fmt.Println("Exit...")
		context.SetState(nil)
	default:
		fmt.Println("Invalid option.")
	}
}

func (s *MainMenuAddUserState) addUser(context *Context) {

}

func (s *MainMenuAddUserState) addHosts(context *Context) {

}
