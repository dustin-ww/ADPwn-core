package service

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/repository"
	"context"

	"ADPwn/database/utils"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService() (*ProjectService, error) {
	db, err := utils.GetDB()
	if err != nil {
		return nil, err
	}

	// Repository initialisieren
	projectRepo := repository.NewSQLProjectRepository(db)
	return &ProjectService{repo: projectRepo}, nil
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetAllProjects(ctx)
}
