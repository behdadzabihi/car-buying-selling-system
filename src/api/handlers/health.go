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

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,helper.GenerateBaseResponse("working",true,0))
	return
}