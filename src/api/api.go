package api

import (
	"fmt"

	"github.com/behdadzabihi/car-buying-selling-system/src/api/middlewares"
	"github.com/behdadzabihi/car-buying-selling-system/src/api/routers"
	"github.com/behdadzabihi/car-buying-selling-system/src/api/validations"
	"github.com/behdadzabihi/car-buying-selling-system/src/config"
	"github.com/behdadzabihi/car-buying-selling-system/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New() //create engin and we use Default() that same do

	RegisterValidators()
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery(),middlewares.LimitByRequest())
	RegisterRoutes(r)
	RegisterSwagger(r,cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterValidators() {
	value, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		value.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine) {
	// r.Use(gin.Logger(),gin.Recovery(),middlewares.TestMiddleware(),middlewares.LimitByRequest())
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		testHandler := v1.Group("/test")
		routers.Health(health)
		routers.TestHandler(testHandler)
	}
}

func RegisterSwagger(r *gin.Engine,cfg *config.Config){
	docs.SwaggerInfo.Title="buying-selling car"
	docs.SwaggerInfo.Description="buying-selling car"
	docs.SwaggerInfo.Version="1.0"
	docs.SwaggerInfo.BasePath="/api"
	docs.SwaggerInfo.Host=fmt.Sprintf("localhost:%s",cfg.Server.Port)
	docs.SwaggerInfo.Schemes=[]string{"http"}

	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}
