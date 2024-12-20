package states

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/service"
	db_context "context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"time"
)

type MainMenuAddHostRange struct {
	App     *tview.Application
	Project model.Project
}

func (s *MainMenuAddHostRange) Execute(context *Context) {
	inputField := tview.NewInputField().
		SetLabel("Add Host Range (Format: 10.10.10.10/24").
		SetFieldWidth(20)

	inputField.SetDoneFunc(func(key tcell.Key) {
		hostRange := inputField.GetText()
		s.addHostRange(s.Project, hostRange)
		context.SetState(&MainMenuState{App: s.App, Project: s.Project})
	})

	s.App.SetRoot(inputField, true).SetFocus(inputField)
}

func (s *MainMenuAddHostRange) addHostRange(project model.Project, hostRange string) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	project, err := projectService.SaveSubnet(ctx, project, hostRange)

	if err != nil {
		log.Fatal("Error while saving new project: ", err)
	}
}
