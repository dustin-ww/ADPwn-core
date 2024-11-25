package states

import (
	"ADPwn/database/project/service"
	db_context "context"
	"fmt"
	"time"
)

type ProjectCreateMenuState struct{}

func (s *ProjectCreateMenuState) Execute(context *Context) {
	fmt.Println("\n Please enter name of project:")
	fmt.Println("1 - Back to main Menu")
	fmt.Println("\n\n\n\n")
	var choice int
	fmt.Scan(&choice)

	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	//ctx, cancel := context.(context.Background(), 5*time.Second)
	defer cancel()

	projectService, _ := service.NewProjectService()
	projects, err := projectService.AllProjects(ctx)

	if err != nil {
		fmt.Printf("Fehler beim Abrufen der Projekte: %v\n", err)
		return
	}
	for _, project := range projects {
		fmt.Printf("Projekt: %s (UUID: %s)\n", project.Name, project.ID)
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
