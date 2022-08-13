package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
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
	id := c.Params.ByName("id")
	nid, _ := strconv.Atoi(id)
	fmt.Println(nid)
	comment := model.ArticleComment{
		Model: gorm.Model{
			ID:        uint(nid),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	a.DB.Preload(clause.Associations).Find(&comment)
	response.Success(c, gin.H{
		"data": comment,
	}, "success")
}

func (a ArticleCommentController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	com := model.ArticleComment{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//if err := c2.DB.Unscoped().Delete(&cat).Error; err != nil {\
	if err := a.DB.Delete(&com).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}

func NewArticleCommentController() IArticleCommentController {
	db := common.GetDb()
	db.AutoMigrate(model.ArticleComment{})

	return ArticleCommentController{db}
}
