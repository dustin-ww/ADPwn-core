package handlers

import (
	"ADPwn/core/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) GetProjectOverviews(c *gin.Context) {
	// Verwende den Service
	projects, err := h.projectService.GetOverviewForAll(c.Request.Context())
	log.Println(*projects[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// Weitere Handler-Methoden hier hinzufügen, z. B.:
func (h *ProjectHandler) AddDomainWithHosts(c *gin.Context) {
	// Implementiere die Logik mit h.projectService
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	type CreateProjectRequest struct {
		Name string `json:"name" binding:"required"`
	}

	var request CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	projectUID, err := h.projectService.Create(
		c.Request.Context(),
		request.Name,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create project",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"uid":     projectUID,
		"message": "Project created successfully",
	})
}
