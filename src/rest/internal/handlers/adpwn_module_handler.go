package handlers

import (
	"ADPwn/core/service"
	"github.com/gin-gonic/gin"
	"net/http"

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
	modules := h.adpwnModuleService.GetAll()
	if len(modules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no adpwn modules found"})
	}
	c.JSON(http.StatusOK, modules)
}
