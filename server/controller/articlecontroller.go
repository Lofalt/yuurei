package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {

	db := common.GetDb()
	article := make([]model.Article, 5)
	fmt.Println(article)
	db.Find(&article)
	response.Success(c, gin.H{"data": article}, "success")

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
