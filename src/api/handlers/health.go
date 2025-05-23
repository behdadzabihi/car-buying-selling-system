package handlers

import (
	"net/http"

	"github.com/behdadzabihi/car-buying-selling-system/src/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,helper.GenerateBaseResponse("working",true,0))
	return
}