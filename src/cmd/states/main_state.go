package states

import (
	"ADPwn/cmd/states/common"
	"ADPwn/database/project/model"
	worker "ADPwn/tools"
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

	// Eine einzige Liste für alle Optionen
	mainMenuList := tview.NewList()

	mainMenuList.AddItem("[green::b] 🎯 "+s.Project.Name+"[-:-:-]", "", 0, nil)

	// Konfigurationsoptionen mit Überschrift hinzufügen
	mainMenuList.AddItem("[yellow::b] ⚙️ Configuration Options[-:-:-]", "", 0, nil)

	mainMenuList.AddItem("Add Single Host", "", '1', func() {
		context.SetState(&AddHostState{App: s.App, Project: s.Project})
	})
	mainMenuList.AddItem("Add Host Range", "", '2', func() {
		context.SetState(&AddHostRangeState{App: s.App, Project: s.Project})
	})
	mainMenuList.AddItem("Add User", "", '3', func() {
		context.SetState(&AddUserState{App: s.App, Project: s.Project})
	})

	// Trennüberschrift hinzufügen
	mainMenuList.AddItem("[red::b] ⚔️ Execution Options[-:-:-]", "", 0, nil)

	// Ausführungsoptionen hinzufügen
	mainMenuList.AddItem("Run Enumeration", "", '4', func() {
		context.SetState(nil)
		s.App.Stop()
		enumerator := worker.Enumerator{}
		enumerator.Run(s.Project)
	})

	mainMenuList.AddItem("Exit", "", 'q', func() {
		context.SetState(nil)
		s.App.Stop()
		os.Exit(0)
	})

	mainMenuList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		// Wenn die Überschrift ausgewählt wird, springe zur nächsten Zeile
		if shortcut == 0 {
			if index < mainMenuList.GetItemCount()-1 {
				mainMenuList.SetCurrentItem(index + 1)
			} else {
				mainMenuList.SetCurrentItem(0)
			}
		}
	})

	// Layout mit der Liste erstellen
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false). // Titel oben
		AddItem(mainMenuList, 0, 1, true) // Eine Liste, die alles enthält

	// Fokus auf die Liste setzen
	s.App.SetRoot(flex, true).SetFocus(mainMenuList)

}
