package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/database/model"
	"ADPwn/database/service"
	db_context "context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"time"
)

type AddSubnetTarget struct {
	App     *tview.Application
	Project model.Project
}

func (s *AddSubnetTarget) Execute(context *common.Context) {
	inputField := tview.NewInputField().
		SetLabel("(Format: 10.10.10.10/24) ").
		SetFieldWidth(20)

	inputField.SetDoneFunc(func(key tcell.Key) {
		hostRange := inputField.GetText()
		s.addHostRange(s.Project, hostRange)
		context.SetState(&MainState{App: s.App, Project: s.Project})
	})

	inputField.SetBorder(true).SetTitle("Add Subnet Target").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(inputField, true).SetFocus(inputField)
}

func (s *AddSubnetTarget) addHostRange(project model.Project, hostRange string) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	project, err := projectService.SaveSubnetTarget(ctx, project, hostRange)

	if err != nil {
		log.Fatal("Error while saving new host range to project: ", err)
	}
}
