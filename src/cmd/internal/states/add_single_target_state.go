package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/database/model"
	"ADPwn/database/service"
	db_context "context"
	"github.com/rivo/tview"
	"time"
)

type AddSingleTargetState struct {
	App     *tview.Application
	Project model.Project

	ip string
}

func (s *AddSingleTargetState) Execute(context *common.Context) {

	form := tview.NewForm().
		AddInputField("IP", "", 20, nil, func(text string) {
			s.ip = text
		}).
		AddButton("Save", func() {
			if err := s.addHost(); err == nil {
				common.ShowSuccessAlert(s.App, "Host added successfully!", func() {
					context.SetState(&MainState{App: s.App, Project: s.Project})
				})
			} else {
				common.ShowErrorAlert(s.App, "Error while adding host!", func() {
					context.SetState(&MainState{App: s.App, Project: s.Project})
				})
			}

		}).
		AddButton("Back", func() {
			context.SetState(&MainState{App: s.App, Project: s.Project})
		})

	form.SetBorder(true).SetTitle("Add Single Host").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(form, true).SetFocus(form)
}

func (s *AddSingleTargetState) addHost() error {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	_, err := projectService.SaveSingleTarget(ctx, s.Project, s.ip)

	return err
}
