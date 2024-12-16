package states

import (
	"ADPwn/database/project/model"
	worker "ADPwn/tools"

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
		context.SetState(nil)
		s.App.Stop()
		enumerator := worker.Enumerator{}
		enumerator.Run()
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false).
		AddItem(list, 0, 1, true)

	s.App.SetRoot(flex, true).SetFocus(list)

}

func (s *MainMenuState) addUser(context *Context) {

}

func (s *MainMenuState) addHosts(context *Context) {

}
