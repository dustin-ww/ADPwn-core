package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/cmd/logger"
	"ADPwn/core/model"
	"ADPwn/modules"
	"ADPwn/modules/enumeration"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

type MainState struct {
	Project model.Project
	App     *tview.Application
	Logger  *logger.ADPwnLogger
}

func (s *MainState) Execute(context *common.Context) {
	s.Logger = logger.NewADPwnLogger()

	title := s.createTitle()
	mainMenuList := s.createMainMenuList(context)
	logView := s.createLogView()

	flex := s.createLayout(title, mainMenuList, logView)

	s.App.SetRoot(flex, true).SetFocus(mainMenuList)
}

func (s *MainState) createTitle() *tview.TextView {
	return tview.NewTextView().
		SetText("ADPwn - Main Menu - " + s.Project.Name).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
}

func (s *MainState) createMainMenuList(context *common.Context) *tview.List {
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

	s.addModulesToMenu(mainMenuList)

	mainMenuList.AddItem("Exit", "", 'q', func() {
		context.SetState(nil)
		s.App.Stop()
		os.Exit(0)
	})

	return mainMenuList
}

func (s *MainState) addModulesToMenu(mainMenuList *tview.List) {
	for _, module := range modules.GetADPwnModules() {
		mod := module.(*enumeration.NetworkExplorer)
		mod.Logger = s.Logger

		mainMenuList.AddItem(
			mod.GetName(),
			"Version: "+mod.GetVersion()+" from: "+mod.GetAuthor(),
			0,
			func() {
				mod.Execute(s.Project, nil)
			},
		)
	}
}

func (s *MainState) createLogView() *tview.TextView {
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			s.App.Draw()
		})
	logView.SetBorder(true).SetTitle(" Live Logs ")

	s.subscribeToLogger(logView)

	return logView
}

func (s *MainState) subscribeToLogger(logView *tview.TextView) {
	logChan := s.Logger.Subscribe()
	go func() {
		for msg := range logChan {
			fmt.Fprintf(logView, "%s\n", msg)
		}
	}()
}

func (s *MainState) createLayout(title *tview.TextView, mainMenuList *tview.List, logView *tview.TextView) *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(title, 3, 0, false).
			AddItem(mainMenuList, 0, 1, true),
			0, 1, true).
		AddItem(logView, 0, 1, false)
}
