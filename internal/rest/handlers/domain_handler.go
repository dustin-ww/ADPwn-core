package handlers

import (
	"ADPwn-core/pkg/service"
)

type DomainHandler struct {
	projectService *service.ProjectService
}

func NewDomainHandler(projectService *service.ProjectService) *DomainHandler {
	return &DomainHandler{
		projectService: projectService,
	}
}
