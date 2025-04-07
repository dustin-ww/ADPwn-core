package handlers

import (
	"ADPwn/core/model"
	"ADPwn/core/service"
	"ADPwn/rest/internal/util"
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
	log.Println(*project)
	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) AddDomainWithHosts(c *gin.Context) {
	panic("implement me")
}

func (h *ProjectHandler) GetTargets(c *gin.Context) {
	uid := c.Param("projectUID")
	if uid == "" {
		fmt.Print("BAD REQUEST")
		fmt.Println(uid)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project UID is required",
		})
		return
	}

	project, err := h.projectService.GetTargets(
		c.Request.Context(),
		uid,
	)

	fmt.Println(err)

	if err != nil {
		errReturn := gin.H{
			"error":   "Failed to retrieve targets",
			"details": err.Error(),
		}

		if err.Error() == "targets not found" {
			c.JSON(http.StatusNotFound, errReturn)
		} else {
			c.JSON(http.StatusInternalServerError, errReturn)
		}
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) CreateTarget(c *gin.Context) {
	type CreateTargetRequest struct {
		IP   string `json:"ip" binding:"required" validate:"required"`
		Note string `json:"note"`
		CIDR int    `json:"cidr"`
	}

	var request CreateTargetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	uid := c.Param("projectUID")
	fmt.Println("UID:", uid)

	target, err := h.projectService.CreateTargets(
		c.Request.Context(),
		uid,
		request.IP,
		request.Note,
		request.CIDR,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create target",
			"details": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, target)
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
	uid := c.Param("projectUID")
	var fields map[string]interface{}

	if err := c.BindJSON(&fields); err != nil {
		c.JSON(400, gin.H{"error": "INVALID_JSON"})
		return
	}
	fmt.Printf("Anzahl der Felder im Handler: %d\n", len(fields))

	if err := h.projectService.UpdateFields(c, uid, fields); err != nil {
		util.HandleServiceError(c, err)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (h *ProjectHandler) AddDomain(c *gin.Context) {
	type AddDomainRequest struct {
		Name string `json:"name" binding:"required" validate:"required"`
	}

	var request AddDomainRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	uid := c.Param("projectUID")

	domain := &model.Domain{
		Name: request.Name,
	}

	err := h.projectService.AddDomain(
		c.Request.Context(),
		uid,
		domain,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to add domain",
			"details": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Domain successfully added",
		"name":   request.Name,
	})
}

func (h *ProjectHandler) GetDomains(c *gin.Context) {
	uid := c.Param("projectUID")
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project UID is required",
		})
		return
	}

	log.Println(uid)
	domains, err := h.projectService.GetProjectDomains(
		c.Request.Context(),
		uid,
	)

	if err != nil {
		errReturn := gin.H{
			"error":   "Failed to retrieve domains",
			"details": err.Error(),
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, errReturn)
	}

	log.Println(domains)

	c.JSON(http.StatusOK, domains)
}
