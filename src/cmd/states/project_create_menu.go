package states

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/service"
	db_context "context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type ProjectCreateMenuState struct{}

func (s *ProjectCreateMenuState) Execute(context *Context) {
	var name string

	fmt.Println("\n Please enter name of project:")
	fmt.Scan(&name)

	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	uuid, _ := uuid.NewV7()
	projectToSave := *model.NewProject(uuid.String(), name)

	projectService, _ := service.NewProjectService()
	err := projectService.SaveProject(ctx, projectToSave)

	if err != nil {
		log.Fatal("Error while creating a new project!")
		os.Exit(1)
	} else {
		println("New project is created")
	}

	context.SetState(&StartMenuState{})
}
