package api

import (
	

	"github.com/behdadzabihi/car-buying-selling-system/api/routers"
	"github.com/gin-gonic/gin"
)


func InitServer(){
	r:= gin.New() //create engin and we use Default() that same do
	r.Use(gin.Logger(),gin.Recovery())

	v1:= r.Group("/api/v1/")
	{
			health := v1.Group("/health")
			routers.Health(health)
		}

		r.Run(":5005")
	}