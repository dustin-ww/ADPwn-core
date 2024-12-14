package states

import (
	"ADPwn/database/project/model"

	"github.com/rivo/tview"
)

type MainMenuState struct {
	Project model.Project
	App     *tview.Application
}

func (s *MainMenuState) Execute(context *Context) {

	title := tview.NewTextView().
		SetText("ADPwn - Main Menu").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	list := tview.NewList()

	list.AddItem("Add Single Host", "", '1', func() {
		context.SetState(&StartMenuState{App: s.App})
	})
	list.AddItem("Add Host Range", "", '2', func() {
		context.SetState(&StartMenuState{App: s.App})
	})
	list.AddItem("Add User", "", '3', func() {
		context.SetState(&StartMenuState{App: s.App})
	})

	list.AddItem("Run Enumeration", "", '4', func() {
		context.SetState(&StartMenuState{App: s.App})
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false).
		AddItem(list, 0, 1, true)

	s.App.SetRoot(flex, true).SetFocus(list)

}

/* func (s *MainMenuState) printCMD() {
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
} */

func (s *MainMenuState) addUser(context *Context) {

}

func (s *MainMenuState) addHosts(context *Context) {

}
