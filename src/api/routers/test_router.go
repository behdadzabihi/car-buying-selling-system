package routers

import (
	"github.com/behdadzabihi/car-buying-selling-system/src/api/handlers"
	"github.com/gin-gonic/gin"
)



func TestHandler(r *gin.RouterGroup){
	testHandler:= handlers.NewTestHandler()
	r.GET("/",testHandler.Test)
	r.GET("/user/:id",testHandler.UserById)
	r.GET("/user/get-user-by-username/:username",testHandler.UserByUsername)
	r.GET("/binder1",testHandler.HeaderBinder1)
	r.GET("/binder2",testHandler.HeaderBinder2)


	r.POST("/binder/query1", testHandler.QueryBinder1)
	r.POST("/binder/query2", testHandler.QueryBinder2)

	r.POST("/binder/uri/:id/:name", testHandler.UriBinder)
	r.POST("/binder/body", testHandler.BodyBinder)
	r.POST("/binder/form", testHandler.FormBinder)
	r.POST("/binder/file", testHandler.FileBinder)

}