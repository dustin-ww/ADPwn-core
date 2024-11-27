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
	fmt.Println("\n\n\n\n")
	var choice int
	fmt.Scan(&choice)

	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	//ctx, cancel := context.(context.Background(), 5*time.Second)
	defer cancel()

	projectService, err := service.NewProjectService()
	projects, err := projectService.AllProjects(ctx)

	if err != nil {
		fmt.Printf("Error while fetching projects: %v\n", err)
		return
	}
	for _, project := range projects {
		fmt.Println("Projekt: " + project.Name)
		fmt.Println("ID: " + fmt.Sprintf("%d", project.ID))
	}

	switch choice {
	case 1:
		context.SetState(&MainMenuState{})
	case 2:
		fmt.Println("Aktion ausgeführt.")
	default:
		fmt.Println("Ungültige Auswahl.")
	}
}
