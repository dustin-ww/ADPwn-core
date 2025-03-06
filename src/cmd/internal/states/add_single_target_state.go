package states

import (
	"ADPwn/cmd/internal/states/common"
	"ADPwn/core/service"
	db_context "context"
	"github.com/rivo/tview"
	"time"
)

type AddSingleTargetState struct {
	App       *tview.Application
	ProjectID string

	ip string
}

func (s *AddSingleTargetState) Execute(context *common.Context) {

	form := tview.NewForm().
		AddInputField("IP", "", 20, nil, func(text string) {
			s.ip = text
		}).
		AddButton("Save", func() {
			if err := s.addTarget(); err == nil {
				common.ShowSuccessAlert(s.App, "Host added successfully!", func() {
					context.SetState(&MainState{App: s.App, ProjectID: s.ProjectID})
				})
			} else {
				common.ShowErrorAlert(s.App, "Error while adding host!", func() {
					context.SetState(&MainState{App: s.App, ProjectID: s.ProjectID})
				})
			}

		}).
		AddButton("Back", func() {
			context.SetState(&MainState{App: s.App, ProjectID: s.ProjectID})
		})

	form.SetBorder(true).SetTitle("Add Single Target").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(form, true).SetFocus(form)
}

func (s *AddSingleTargetState) addTarget() error {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, _ := service.NewProjectService()
	err := projectService.AddTarget(ctx, s.ProjectID, s.ip)
	return err
}
