package states

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/service"
	db_context "context"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectCreateMenuState struct {
	App *tview.Application
}

func (s *ProjectCreateMenuState) Execute(context *Context) {
	inputField := tview.NewInputField().
		SetLabel("Enter a project name: ").
		SetFieldWidth(20)

	inputField.SetDoneFunc(func(key tcell.Key) {
		projectName := inputField.GetText()
		s.createProject(projectName)

		context.SetState(&StartMenuState{App: s.App})

	})

	s.App.SetRoot(inputField, true).SetFocus(inputField)
}

func (s *ProjectCreateMenuState) createProject(projectName string) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectToSave := *model.NewProject(projectName)

	projectService, _ := service.NewProjectService()
	err := projectService.SaveProject(ctx, projectToSave)

	if err != nil {
		log.Fatal("Error while saving new project: ", err)
		panic("Failed to save project")
	}
}
