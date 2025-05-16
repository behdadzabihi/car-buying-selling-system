package api

import (
	"fmt"

	"github.com/behdadzabihi/car-buying-selling-system/api/routers"
	"github.com/behdadzabihi/car-buying-selling-system/api/validations"
	"github.com/behdadzabihi/car-buying-selling-system/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


func InitServer(){
	cfg:= config.GetConfig()
	r:= gin.New() //create engin and we use Default() that same do

	value,ok:=binding.Validator.Engine().(*validator.Validate)
	if ok{
		value.RegisterValidation("mobile",validations.IranianMobileNumberValidation,true)
	}

	
	r.Use(gin.Logger(),gin.Recovery())
	api:=r.Group("/api")
	v1:= api.Group("/v1")
	{
			health := v1.Group("/health")
			testHandler:=v1.Group("/test")
			routers.Health(health)
			routers.TestHandler(testHandler)
	}
		r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	}