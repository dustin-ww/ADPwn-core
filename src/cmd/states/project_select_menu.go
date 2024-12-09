package states

import (
	"ADPwn/database/project/service"
	db_context "context"
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

type ProjectSelectMenuState struct{}

func (s *ProjectSelectMenuState) Execute(context *Context) {
	tm.Clear()
	tm.Flush()

	fmt.Println("\nPlease Select a project to load:")
	fmt.Println("1 - Back to main Menu")

	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	projectService, serviceErr := service.NewProjectService()
	projects, err := projectService.AllProjects(ctx)

	if serviceErr != nil {
		fmt.Printf("Error while fetching projects: %v\n", err)
		return
	}

	for i, project := range projects {
		fmt.Println(i, " -----------")
		fmt.Println("Projekt: " + project.Name)
		fmt.Println("ID: " + project.UID + "\n")
		fmt.Println("UID: " + project.UID + "\n")
	}

	fmt.Println("Enter Project Number: ")

	var choice int
	fmt.Scan(&choice)

	if choice >= len(projects) {
		fmt.Println("Error: your input is invalid!")
	} else {
		context.SetState(&MainMenuState{projects[choice-1]})
	}

}
