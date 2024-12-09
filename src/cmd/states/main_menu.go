package states

import (
	"ADPwn/cmd/cmdutils"
	"ADPwn/database/project/model"
	"fmt"

	tm "github.com/buger/goterm"
)

type MainMenuState struct {
	Project model.Project
}

func (s *MainMenuState) Execute(context *Context) {
	cmdutils.ClearCMD()
	s.printCMD()

	var choice int
	fmt.Print("\n Please choose options: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		context.SetState(&MainMenuAddUserState{Project: s.Project})
	case 2:
		fmt.Println("Exit...")
		context.SetState(nil)
	default:
		fmt.Println("Invalid option.")
	}
}

func (s *MainMenuState) printCMD() {
	tm.Println(tm.Background(tm.Color(tm.Bold("ADPwn - Main Menu"), tm.RED), tm.WHITE))
	tm.Println(tm.Background(tm.Color(tm.Bold("\nProject: "+s.Project.Name+"("+s.Project.UID+")"), tm.RED), tm.WHITE))

	tm.Println(tm.Color(tm.Bold("--- Configure Project ---"), tm.YELLOW))
	tm.Println(tm.Color(tm.Bold("1. Add Hosts"), tm.YELLOW))
	tm.Println(tm.Color(tm.Bold("2. Add User"), tm.YELLOW))

	tm.Println(tm.Color(tm.Bold("--- Get Informations ---"), tm.BLUE))
	tm.Println(tm.Color(tm.Bold("3. Get all"), tm.BLUE))

	tm.Println(tm.Color(tm.Bold("--- Run Attacks ---"), tm.RED))
	tm.Println(tm.Color(tm.Bold("4. Run nmap enumeration"), tm.RED))

	tm.Flush()
}

func (s *MainMenuState) addUser(context *Context) {

}

func (s *MainMenuState) addHosts(context *Context) {

}
