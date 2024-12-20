package states

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/service"
	db_context "context"
	"fmt"
	"log"
	"time"

	"github.com/rivo/tview"
)

type ProjectSelectMenuState struct {
	App *tview.Application
}

func (s *ProjectSelectMenuState) Execute(context *Context) {
	title := s.createTitle("ADPwn - Select Project")

	list := tview.NewList()
	list.AddItem("1. Back to Main Menu", "Go back to the main menu", '1', func() {
		context.SetState(&MainMenuAddHostRange{App: s.App})
	})

	for i, project := range s.fetchProjects() {
		index := i + 1
		list.AddItem(
			fmt.Sprintf("%d. %s", index, project.Name),
			"",
			rune('1'+i),
			func() { s.showActions(*context, project) },
		)
	}

	s.setRootLayout(title, list)
}

func (s *ProjectSelectMenuState) showActions(context Context, project model.Project) {
	title := s.createTitle("Project Actions")
	list := tview.NewList().
		AddItem("1. Load", "Load Project", '1', func() {
			s.loadProject(context, project)
		}).
		AddItem("2. Delete", "Delete Project", '2', func() {
			s.deleteProject(context, project)
		})

	s.setRootLayout(title, list)
}

func (s *ProjectSelectMenuState) loadProject(context Context, project model.Project) {
	context.SetState(&MainMenuState{Project: project, App: s.App})
}

func (s *ProjectSelectMenuState) deleteProject(context Context, project model.Project) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)
	defer cancel()

	if projectService, err := service.NewProjectService(); err == nil {
		projectService.DeleteProject(ctx, project)
		context.SetState(&StartMenuState{App: s.App})
	} else {
		log.Printf("Error deleting project: %v", err)
	}
}

func (s *ProjectSelectMenuState) fetchProjects() []model.Project {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)
	defer cancel()

	projectService, err := service.NewProjectService()
	if err != nil {
		log.Fatal("Error creating project service: ", err)
	}

	projects, err := projectService.AllProjects(ctx)
	if err != nil {
		log.Fatal("Error fetching projects: ", err)
	}

	//json, _ := json.MarshalIndent(projects, "", "  ")

	return projects
}

func (s *ProjectSelectMenuState) createTitle(text string) *tview.TextView {
	return tview.NewTextView().
		SetText(text).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
}

func (s *ProjectSelectMenuState) setRootLayout(title *tview.TextView, content tview.Primitive) {
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 3, 0, false).
		AddItem(content, 0, 1, true)

	s.App.SetRoot(flex, true).SetFocus(content)
}
