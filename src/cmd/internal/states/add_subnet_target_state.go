package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/core/service"
	db_context "context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"time"
)

type AddSubnetTarget struct {
	App       *tview.Application
	ProjectID string
}

func (s *AddSubnetTarget) Execute(context *common.Context) {
	inputField := tview.NewInputField().
		SetLabel("(Format: 10.10.10.10/24) ").
		SetFieldWidth(20)

	inputField.SetDoneFunc(func(key tcell.Key) {
		hostRange := inputField.GetText()
		s.addHostRange(hostRange)
		context.SetState(&MainState{App: s.App, ProjectID: s.ProjectID})
	})

	inputField.SetBorder(true).SetTitle("Add Subnet Target").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(inputField, true).SetFocus(inputField)
}

func (s *AddSubnetTarget) addHostRange(hostRange string) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	err := projectService.AddTarget(ctx, projectService, hostRange)

	if err != nil {
		log.Fatal("Error while saving new host range to project: ", err)
	}
}
