package states

import (
	"ADPwn/database/project/service"
	db_context "context"
	"fmt"
	"time"
)

type ProjectSelectMenuState struct{}

func (s *ProjectSelectMenuState) Execute(context *Context) {
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
		fmt.Println("ID: " + project.ID + "\n")
	}

	fmt.Println("Enter Project Number: ")

	var choice int
	fmt.Scan(&choice)

	context.SetState(&MainMenuState{projects[choice]})

}
