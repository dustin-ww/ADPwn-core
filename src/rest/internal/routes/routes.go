package routes

import (
	"ADPwn/core/service"
	"ADPwn/rest/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProjectHandlers(router *gin.Engine, projectService *service.ProjectService) {
	handler := handlers.NewProjectHandler(projectService)

	projectGroup := router.Group("/projects")
	{
		projectGroup.GET("/overviews", handler.GetProjectOverviews)
		projectGroup.GET("/all", handler.GetProjectOverviews)
		projectGroup.GET("/:UID", handler.Get)
		projectGroup.PATCH("/:UID", handler.UpdateProject)
		projectGroup.POST("/", handler.CreateProject)
		projectGroup.POST("/:UID/domains", handler.AddDomainWithHosts)
		projectGroup.POST("/:UID/hk", handler.AddDomainWithHosts)
		projectGroup.GET("/:UID/targets", handler.GetTargets)
		projectGroup.POST("/:UID/targets", handler.CreateTarget)
	}
}

func RegisterADPwnModuleHandlers(router *gin.Engine, adpwnmModuleService *service.ADPwnModuleService) {
	handler := handlers.NewADPwnModuleHandler(adpwnmModuleService)

	projectGroup := router.Group("/adpmod")
	{
		projectGroup.GET("/", handler.GetModules)
		projectGroup.GET("/graph", handler.GetModuleInheritanceGraph)

	}
}
