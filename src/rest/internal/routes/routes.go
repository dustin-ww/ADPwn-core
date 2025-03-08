package routes

import (
	"ADPwn/core/service"
	"ADPwn/rest/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProjectHandlers(router *gin.Engine, projectService *service.ProjectService) {
	handler := handlers.NewProjectHandler(projectService)

	// Definiere die Routen
	projectGroup := router.Group("/projects")
	{
		projectGroup.GET("/", handler.GetProjectOverviews)
		projectGroup.POST("/", handler.CreateProject)
		projectGroup.POST("/:projectUID/domains", handler.AddDomainWithHosts)
		projectGroup.POST("/:projectUID/targets", handler.AddDomainWithHosts)
	}
}
