package routes

import (
	"ADPwn/core/service"
	"ADPwn/rest/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProjectHandlers(router *gin.Engine, projectService *service.ProjectService) {
	projectHandler := handlers.NewProjectHandler(projectService)

	projectGroup := router.Group("/projects")
	{
		projectGroup.GET("/overviews", projectHandler.GetProjectOverviews)
		projectGroup.GET("/", projectHandler.GetProjectOverviews)
		projectGroup.POST("/", projectHandler.CreateProject)

		projectItemGroup := projectGroup.Group("/:projectUID")
		{
			projectItemGroup.GET("", projectHandler.Get)
			projectItemGroup.PATCH("", projectHandler.UpdateProject)

			domainsGroup := projectItemGroup.Group("/domains")
			{
				domainsGroup.GET("", projectHandler.GetDomains)
				domainsGroup.POST("", projectHandler.AddDomain)
			}
			targetsGroup := projectItemGroup.Group("/targets")
			{
				targetsGroup.GET("", projectHandler.GetTargets)
				targetsGroup.POST("", projectHandler.CreateTarget)
			}
		}
	}
}

func RegisterADPwnModuleHandlers(router *gin.Engine, adpwnModuleService *service.ADPwnModuleService) {
	moduleHandler := handlers.NewADPwnModuleHandler(adpwnModuleService)

	moduleGroup := router.Group("/adpwn")
	{
		moduleGroup.GET("/", moduleHandler.GetModules)
		moduleGroup.GET("/graph", moduleHandler.GetModuleInheritanceGraph)
	}
}
