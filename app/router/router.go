package router

import (
	"denny/gin_gorm_practise/app/controller"

	"github.com/gin-gonic/gin"
)

type ControllerList struct {
	controller.IndexController
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	c := ControllerList{}

	r.GET("/", c.IndexController.Index)

	return r
}
