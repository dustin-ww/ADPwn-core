package service

import (
	"ADPwn/database/project/model"
	"ADPwn/database/project/repository"
)

// ProjectService enthält die Geschäftslogik für die `Project`-Entität.
type ProjectService struct {
	repo repository.ProjectRepository
}

// NewProjectService erstellt eine neue Instanz von ProjectService.
func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// GetAllProjects ruft alle Projekte über das Repository ab.
func (s *ProjectService) GetAllProjects() ([]model.Project, error) {
	return s.repo.GetAllProjects()
}
