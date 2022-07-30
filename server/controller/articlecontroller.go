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

type IArticleController interface {
	RESTfulcontroller
	SelectAll(c *gin.Context)
	SelectArticleById(c *gin.Context)
	SelectArticlesByTagId(c *gin.Context)
	SelectArticlesByCategoryId(c *gin.Context)
}

type ArticleController struct {
	DB *gorm.DB
}

func (a ArticleController) SelectAll(c *gin.Context) {
	var articles []model.Article

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	a.DB.Model(model.Article{}).Count(&total)
	if err := a.DB.Preload(clause.Associations).Offset(offset).Limit(pageSize).Find(&articles).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{"data": articles, "total": total}, "success")
}

func (a ArticleController) SelectArticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	nid, _ := strconv.Atoi(id)
	fmt.Println(nid)
	article := model.Article{
		Model: gorm.Model{
			ID:        uint(nid),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	a.DB.Debug().Preload("Tags").Preload("ArticleCategory").Find(&article)
	response.Success(c, gin.H{
		"data": article,
	}, "success")

}

func (a ArticleController) SelectArticlesByTagId(c *gin.Context) {
	tagId, _ := strconv.Atoi(c.Params.ByName("id"))
	Id := uint(tagId)
	//var articles []*model.Article
	tag := model.Tag{
		Model: gorm.Model{
			ID: Id,
		},
	}
	//pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	//pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	//var offset int
	//if pageNum == -1 {
	//	offset = -1
	//} else {
	//	offset = (pageNum - 1) * pageSize
	//}
	//var total int64
	//a.DB.Model(model.Article{}).Where(&tag).Count(&total)
	//db.Model(&tag).Association("Articles")

	//err1 := db.Model(&tag).Association("Articles").Find(&articles)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	a.DB.Debug().Preload("Articles").Find(&tag)
	response.Success(c, gin.H{"data": tag}, "success")
}

func (a ArticleController) SelectArticlesByCategoryId(c *gin.Context) {
	id := c.Params.ByName("id")
	tid, _ := strconv.Atoi(id)
	article := model.Article{
		ArticleTitle:      "",
		ArticleContent:    "",
		ArticleReadTimes:  0,
		ArticleLoveTimes:  0,
		IsActive:          false,
		ArticleSummary:    "",
		ArticleCategoryID: tid,
		ArticleCategory:   &model.ArticleCategory{},
		Tags:              nil,
	}
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	a.DB.Model(model.Article{}).Where(&article).Count(&total)
	var articles []model.Article
	a.DB.Debug().Preload(clause.Associations).Where(&article).Offset(offset).Limit(pageSize).Find(&articles)
	response.Success(c, gin.H{"data": articles, "total": total}, "success")
}

func CreateArticleController() IArticleController {
	db := common.GetDb()
	err := db.Debug().AutoMigrate(model.Article{}, model.Tag{}, model.ArticleComment{})
	if err != nil {
		fmt.Println(err.Error())

	}
	return ArticleController{DB: db}
}
func (a ArticleController) Create(c *gin.Context) {
	var article model.Article
	err := c.ShouldBind(&article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(article)
	result := a.DB.Debug().Create(&article)
	if result.Error != nil {
		response.Fail(c, gin.H{"错误信息": result.Error}, "出错了")
	} else {
		response.Success(c, gin.H{"data": article}, "添加成功")
	}
}

func (a ArticleController) Update(c *gin.Context) {
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var art model.Article

	var newArt model.Article
	err3 := c.ShouldBind(&newArt)
	if err3 != nil {
		response.Fail(c, gin.H{}, err3.Error())
		return
	}

	if err := a.DB.Where("id=?", catid).First(&art).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := a.DB.Model(&art).Updates(&newArt).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": art}, "更改成功")
}

func (a ArticleController) Show(c *gin.Context) {
	id := c.PostForm("ID")
	nid, _ := strconv.Atoi(id)
	fmt.Println(nid)
	article := model.Article{
		Model: gorm.Model{
			ID:        uint(nid),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	if err := a.DB.Debug().Preload("Tags").Preload("ArticleCategory").Find(&article).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{
		"data": article,
	}, "success")
}

func (a ArticleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	art := model.Article{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//if err := c2.DB.Debug().Unscoped().Delete(&cat).Error; err != nil {\
	if err := a.DB.Debug().Delete(&art).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}
