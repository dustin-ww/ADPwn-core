package rest

import (
	"ADPwn-core/internal/rest/handlers"
	"ADPwn-core/pkg/service"
	"github.com/gin-gonic/gin"
)

func RegisterProjectHandlers(router *gin.Engine, projectService *service.ProjectService) {
	projectHandler := handlers.NewProjectHandler(projectService)

	projectGroup := router.Group("/project")
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

	adpwnGroup := router.Group("/adpwn")
	{
		moduleGroup := adpwnGroup.Group("/modulelib")
		{
			moduleGroup.GET("/", moduleHandler.GetModules)
			moduleGroup.GET("/graph", moduleHandler.GetModuleInheritanceGraph)

			moduleItemGroup := moduleGroup.Group("/:moduleKey")
			{
				moduleItemGroup.GET("/run", moduleHandler.RunModule)
				moduleItemGroup.GET("/options", moduleHandler.GetModuleOptions)

				moduleItemVectorGroup := moduleItemGroup.Group("/vector")
				{
					moduleItemVectorGroup.GET("/run", moduleHandler.RunAttackVector)
					moduleItemVectorGroup.GET("/options", moduleHandler.GetAttackVectorOptions)
				}
			}

		}

	}
}
