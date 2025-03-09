package handlers

import (
	"ADPwn/core/service"
	"fmt"
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
	projects, err := h.projectService.GetOverviewForAll(c.Request.Context())
	log.Println(*projects[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) Get(c *gin.Context) {
	uid := c.Param("projectUID")
	if uid == "" {
		fmt.Print("BAD REQUEST")
		fmt.Println(uid)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project UID is required",
		})
		return
	}

	project, err := h.projectService.Get(
		c.Request.Context(),
		uid,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve project",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) AddDomainWithHosts(c *gin.Context) {
	panic("implement me")
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

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	panic("implement me")
}
