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

func SelectAll(c *gin.Context) {
	db := common.GetDb()
	//article := make([]model.Article, 5)
	//fmt.Println(article)
	//db.Find(&article)
	//response.Success(c, gin.H{"data": article}, "success")
	//article := model.Article{
	//	Model:gorm.Model{
	//		ID:        0,
	//		CreatedAt: time.Time{},
	//		UpdatedAt: time.Time{},
	//		DeletedAt: gorm.DeletedAt{},
	//	},
	//	ArticleTitle:         "",
	//	ArticleContent:       "",
	//	ArticleReadTimes:     0,
	//	ArticleLoveTimes:     0,
	//	IsActive:             false,
	//	ArticleSummary:       "",
	//	ArticleCategoryRefer: 0,
	//	ArticleCategory:      model.ArticleCategory{},
	//	Tags:                 nil,
	//}
	var articles []model.Article
	db.Preload(clause.Associations).Find(&articles)
	response.Success(c, gin.H{"data": articles}, "success")
}

func SelectById(c *gin.Context) {
	id := c.PostForm("articleId")
	db := common.GetDb()
	article := model.Article{}
	db.First(&article, "article_id = ?", id)
	response.Success(c, gin.H{
		"data": article,
	}, "success")
}

func SelectArticleById(c *gin.Context) {
	db := common.GetDb()
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
	//var tags []*model.Tag
	//开启关联查询
	//db.Debug().Model(&article).Association("Tags")
	//err1 := db.Model(&article).Association("Tags").Find(&tags)
	//db.Debug().Where("article_id=?", id).First(&article)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//for _, tag := range tags {
	//	fmt.Println(tag)
	//}
	//article.Tags = tags
	db.Debug().Preload("Tags").Preload("ArticleCategory").Find(&article)
	response.Success(c, gin.H{
		"data": article,
	}, "success")
}

func SelectArticlesByTagId(c *gin.Context) {
	db := common.GetDb()
	tagId, _ := strconv.Atoi(c.PostForm("ID"))
	Id := uint(tagId)
	//var articles []*model.Article
	tag := model.Tag{
		Model: gorm.Model{
			ID: Id,
		},
	}
	//db.Model(&tag).Association("Articles")

	//err1 := db.Model(&tag).Association("Articles").Find(&articles)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	db.Debug().Preload("Articles").Find(&tag)
	response.Success(c, gin.H{"data": tag}, "success")
}

func SelectArticlesByCategoryId(c *gin.Context) {
	db := common.GetDb()
	id := c.PostForm("ID")
	tid, _ := strconv.Atoi(id)
	article := model.Article{
		Model: gorm.Model{
			ID:        uint(tid),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		ArticleTitle:         "",
		ArticleContent:       "",
		ArticleReadTimes:     0,
		ArticleLoveTimes:     0,
		IsActive:             false,
		ArticleSummary:       "",
		ArticleCategoryRefer: tid,
		ArticleCategory:      model.ArticleCategory{},
		Tags:                 nil,
	}
	var articles []model.Article
	db.Preload(clause.Associations).Where(&article).Find(&articles)
	response.Success(c, gin.H{"data": articles}, "success")
}

func AddArticle(c *gin.Context) {
	db := common.GetDb()
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(article)
	result := db.Debug().Create(&article)
	fmt.Println(result)
}
