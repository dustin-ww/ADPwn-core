package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/database/model"
	"ADPwn/modules"
	"os"

	"github.com/rivo/tview"
)

type MainState struct {
	Project model.Project
	App     *tview.Application
}

func (s *MainState) Execute(context *common.Context) {

	title := tview.NewTextView().
		SetText("ADPwn - Main Menu - " + s.Project.Name).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	mainMenuList := tview.NewList()

	mainMenuList.AddItem("[green::b] 🎯 "+s.Project.Name+"[-:-:-]", "", 0, nil)

	mainMenuList.AddItem("[yellow::b] ⚙️ Configuration Options[-:-:-]", "", 0, nil)

	mainMenuList.AddItem("Add Single Host", "", '1', func() {
		context.SetState(&AddSingleTargetState{App: s.App, Project: s.Project})
	})
	mainMenuList.AddItem("Add Host Range", "", '2', func() {
		context.SetState(&AddSubnetTarget{App: s.App, Project: s.Project})
	})
	mainMenuList.AddItem("Add User", "", '3', func() {
		context.SetState(&AddUserState{App: s.App, Project: s.Project})
	})

	mainMenuList.AddItem("[red::b] ⚔️ Execution Options[-:-:-]", "", 0, nil)

	for _, module := range modules.GetADPwnModules() {
		mainMenuList.AddItem(module.GetName(), "Version: "+module.GetVersion()+" from: "+module.GetAuthor(), 0, nil)
	}

	mainMenuList.AddItem("Exit", "", 'q', func() {
		context.SetState(nil)
		s.App.Stop()
		os.Exit(0)
	})

	mainMenuList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if shortcut == 0 {
			if index < mainMenuList.GetItemCount()-1 {
				mainMenuList.SetCurrentItem(index + 1)
			} else {
				mainMenuList.SetCurrentItem(0)
			}
		}
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false). // Titel oben
		AddItem(mainMenuList, 0, 1, true) // Eine Liste, die alles enthält

	s.App.SetRoot(flex, true).SetFocus(mainMenuList)

}
