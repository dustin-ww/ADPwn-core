package handlers

import (
	"ADPwn/core/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "ADPwn/modules/attacks"
	_ "ADPwn/modules/enumeration"
)

type ADPwnModuleHandler struct {
	adpwnModuleService *service.ADPwnModuleService
}

func NewADPwnModuleHandler(adpwnModuleService *service.ADPwnModuleService) *ADPwnModuleHandler {
	return &ADPwnModuleHandler{
		adpwnModuleService: adpwnModuleService,
	}
}

func (h *ADPwnModuleHandler) GetModules(c *gin.Context) {
	modules, err := h.adpwnModuleService.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("failed to get all modules: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(modules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no adpwn modules found"})
	}
	c.JSON(http.StatusOK, modules)
}

func (h *ADPwnModuleHandler) GetModuleInheritanceGraph(c *gin.Context) {
	graph, err := h.adpwnModuleService.GetInheritanceGraph(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, graph)
}

func (h *ADPwnModuleHandler) RunModule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (h *ADPwnModuleHandler) RunAttackVector(c *gin.Context) {
	log.Println("RUN")
	moduleKey := c.Param("moduleKey")
	err := h.adpwnModuleService.RunAttackVector(c.Request.Context(), moduleKey)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
