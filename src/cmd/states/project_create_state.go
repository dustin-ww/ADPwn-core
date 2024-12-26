package states

import (
	"ADPwn/cmd/states/common"
	"ADPwn/database/project/service"
	db_context "context"
	"time"

	"github.com/rivo/tview"
)

type ProjectCreateState struct {
	App *tview.Application

	projectName string
}

func (s *ProjectCreateState) Execute(context *common.Context) {

	form := tview.NewForm().
		AddInputField("Project Name", "", 20, nil, func(text string) {
			s.projectName = text
		}).
		AddButton("Save", func() {
			if err := s.createProject(); err == nil {
				common.ShowSuccessAlert(s.App, "Project created successfully!", func() {
					context.SetState(&StartState{App: s.App})
				})
			} else {
				common.ShowErrorAlert(s.App, "Error while creating project!", func() {
					context.SetState(&StartState{App: s.App})
				})
			}

		}).
		AddButton("Back", func() {
			context.SetState(&StartState{App: s.App})
		})

	form.SetBorder(true).SetTitle("Create new Project").SetTitleAlign(tview.AlignLeft)

	s.App.SetRoot(form, true).SetFocus(form)
}

func (s *ProjectCreateState) createProject() error {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()
	projectService, _ := service.NewProjectService()

	_, err := projectService.CreateProject(ctx, s.projectName)

	return err
}
