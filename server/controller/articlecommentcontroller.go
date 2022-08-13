package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IArticleCommentController interface {
	RESTfulcontroller
}

type ArticleCommentController struct {
	DB *gorm.DB
}

func (a ArticleCommentController) Create(c *gin.Context) {
	var comment model.ArticleComment
	err := c.ShouldBind(&comment)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comment)
	result := a.DB.Create(&comment)
	if result.Error != nil {
		response.Fail(c, gin.H{"错误信息": result.Error}, "出错了")
	} else {
		response.Success(c, gin.H{"data": comment}, "添加成功")
	}
}

func (a ArticleCommentController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleCommentController) Show(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleCommentController) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewArticleCommentController() IArticleCommentController {
	db := common.GetDb()
	db.AutoMigrate(model.ArticleComment{})

	return ArticleCommentController{db}
}
