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

	articleRoutes := r.Group("/yuurei/article")
	artilceController := controller.CreateArticleController()
	articleRoutes.GET("all", artilceController.SelectAll)
	articleRoutes.GET(":id", artilceController.SelectArticleById)
	//r.GET("/yuurei/article", controller.SelectArticleById)
	articleRoutes.GET("tagId/:id", artilceController.SelectArticlesByTagId)
	articleRoutes.GET("categoryId/:id", artilceController.SelectArticlesByCategoryId)
	articleRoutes.POST("", artilceController.Create)
	articleRoutes.PUT(":id", artilceController.Update)
	articleRoutes.DELETE(":id", artilceController.Delete)
	articleRoutes.POST("pre", artilceController.SelectPre)

	categoryRoutes := r.Group("/yuurei/articleCategory")
	categoryController := controller.CreateCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT(":id", categoryController.Update)
	categoryRoutes.DELETE(":id", categoryController.Delete)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.GET("all", categoryController.ShowAll)

	tagRoutes := r.Group("/yuurei/tag")
	tagController := controller.NewTagController()
	tagRoutes.POST("", tagController.Create)
	tagRoutes.PUT(":id", tagController.Update)
	tagRoutes.DELETE(":id", tagController.Delete)
	tagRoutes.GET(":name", tagController.Show)
	tagRoutes.GET("all", tagController.ShowAll)

	msgRoutes := r.Group("/yuurei/msg")
	msgController := controller.NewMessageController()
	msgRoutes.POST("", msgController.Create)
	msgRoutes.PUT(":id", msgController.Update)
	msgRoutes.DELETE(":id", msgController.Delete)
	msgRoutes.GET(":name", msgController.Show)
	msgRoutes.GET("all", msgController.ShowAll)

	picRoutes := r.Group("/yuurei/gallery")
	picController := controller.NewIPicController()
	picRoutes.POST("", picController.Create)
	picRoutes.GET("all", picController.ShowAll)
	picRoutes.DELETE(":id", picController.Delete)
	picRoutes.PUT(":id", picController.Update)

	r.POST("/yuurei/uploadImg", controller.UploadImg)
	r.POST("/yuurei/uploadGallery", controller.GalleryPic)
	r.POST("/yuurei/cropPic", controller.CropImg)

	return r
}
