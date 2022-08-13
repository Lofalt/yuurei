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
	r.POST("/yuurei/validEmail", controller.SendEmail)

	articleRoutes := r.Group("/yuurei/article")
	artilceController := controller.CreateArticleController()
	articleRoutes.GET("all", artilceController.SelectAll)
	articleRoutes.GET(":id", artilceController.SelectArticleById)
	//r.GET("/yuurei/article", controller.SelectArticleById)
	articleRoutes.GET("tagId/:id", artilceController.SelectArticlesByTagId)
	articleRoutes.GET("categoryId/:id", artilceController.SelectArticlesByCategoryId)
	articleRoutes.POST("", middleware.Inception(), artilceController.Create)
	articleRoutes.PUT(":id", middleware.Inception(), artilceController.Update)
	articleRoutes.DELETE(":id", middleware.Inception(), artilceController.Delete)
	articleRoutes.POST("pre", artilceController.SelectPre)

	categoryRoutes := r.Group("/yuurei/articleCategory")
	categoryController := controller.CreateCategoryController()
	categoryRoutes.POST("", middleware.Inception(), categoryController.Create)
	categoryRoutes.PUT(":id", middleware.Inception(), categoryController.Update)
	categoryRoutes.DELETE(":id", middleware.Inception(), categoryController.Delete)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.GET("all", categoryController.ShowAll)

	tagRoutes := r.Group("/yuurei/tag")
	tagController := controller.NewTagController()
	tagRoutes.POST("", middleware.Inception(), tagController.Create)
	tagRoutes.PUT(":id", middleware.Inception(), tagController.Update)
	tagRoutes.DELETE(":id", middleware.Inception(), tagController.Delete)
	tagRoutes.GET(":name", tagController.Show)
	tagRoutes.GET("all", tagController.ShowAll)

	msgRoutes := r.Group("/yuurei/msg")
	msgController := controller.NewMessageController()
	msgRoutes.POST("", msgController.Create)
	msgRoutes.PUT(":id", middleware.Inception(), msgController.Update)
	msgRoutes.DELETE(":id", middleware.Inception(), msgController.Delete)
	msgRoutes.GET(":id", msgController.Show)
	msgRoutes.GET("all", msgController.ShowAll)

	picRoutes := r.Group("/yuurei/gallery")
	picController := controller.NewIPicController()
	picRoutes.POST("", picController.Create)
	picRoutes.GET("all", picController.ShowAll)
	picRoutes.DELETE(":id", picController.Delete)
	picRoutes.PUT(":id", picController.Update)

	setRoutes := r.Group("/yuurei/settings")
	setController := controller.NewSettingController()
	setRoutes.PUT(":id", middleware.Inception(), setController.Update)
	setRoutes.GET(":id", setController.Show)
	setRoutes.POST("", middleware.Inception(), setController.Create)

	entryRoutes := r.Group("/yuurei/entry")
	entryController := controller.NewEntryController()
	entryRoutes.POST("", middleware.Inception(), entryController.Create)
	entryRoutes.GET(":id", entryController.GetEntryById)
	entryRoutes.GET("all", entryController.ShowAll)
	entryRoutes.DELETE(":id", middleware.Inception(), entryController.Delete)
	entryRoutes.GET("categoryId/:id", entryController.GetEntriesByCategoryID)
	entryRoutes.PUT(":id", middleware.Inception(), entryController.Update)

	entryCategoryRoutes := r.Group("/yuurei/entryCategory")
	entryCategoryController := controller.NewEntryCategoryController()
	entryCategoryRoutes.GET("all", entryCategoryController.ShowAll)
	entryCategoryRoutes.GET(":id", entryCategoryController.Show)
	entryCategoryRoutes.POST("", middleware.Inception(), entryCategoryController.Create)
	entryCategoryRoutes.PUT(":id", middleware.Inception(), entryCategoryController.Update)
	entryCategoryRoutes.DELETE(":id", middleware.Inception(), entryController.Delete)

	articleCommentsRoutes := r.Group("/yuurei/articleComment")
	articleCommentController := controller.NewArticleCommentController()
	articleCommentsRoutes.POST("", articleCommentController.Create)
	articleCommentsRoutes.GET(":id", articleCommentController.Show)
	articleCommentsRoutes.DELETE(":id", articleCommentController.Delete)

	r.POST("/yuurei/uploadImgArticle", controller.ArticleImg)
	r.POST("/yuurei/uploadImg", controller.UploadImg)
	r.POST("/yuurei/uploadGallery", controller.GalleryPic)
	r.POST("/yuurei/cropPic", controller.CropImg)
	r.DELETE("/yuurei/file", controller.DeleteFile)

	return r
}
