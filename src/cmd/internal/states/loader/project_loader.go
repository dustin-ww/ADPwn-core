package loader

import (
	"ADPwn/core/model"
	"ADPwn/core/service"
	db_context "context"
	"log"
	"time"
)

func LoadProjectFromDB(projectID string) (model.Project, error) {
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)
	defer cancel()

	projectService, err := service.NewProjectService()
	if err != nil {
		log.Fatal("Error creating project service: ", err)
	}

	project, err := projectService.Get(ctx, projectID)
	if err != nil {
		log.Fatal("Error fetching projects: ", err)
	}

	return *project, err
}
