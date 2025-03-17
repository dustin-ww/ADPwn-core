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

func NewADPwnModuleHandler(adpwnModuleServic *service.ADPwnModuleService) *ADPwnModuleHandler {
	return &ADPwnModuleHandler{
		adpwnModuleService: adpwnModuleServic,
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
