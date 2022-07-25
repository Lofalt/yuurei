package main

import (
	"github.com/Lofalt/yuurei/controller"
	"github.com/Lofalt/yuurei/middleware"
	"github.com/gin-gonic/gin"
)

func CollecRouter(r *gin.Engine) *gin.Engine {

	r.POST("/yuurei/sign_up", controller.Register)
	r.POST("/yuurei/login", controller.Login)
	r.GET("/yuurei/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/yuurei/img", controller.Imgs)
	r.GET("/yuurei/articles", controller.SelectAll)
	r.POST("/yuurei/article", controller.SelectById)
	return r
}