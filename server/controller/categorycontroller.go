package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type CategoryController struct {
	DB *gorm.DB
}

func CreateCategoryController() ICategoryController {
	db := common.GetDb()
	db.AutoMigrate(model.ArticleCategory{})
	return CategoryController{DB: db}
}

func (c2 CategoryController) Create(c *gin.Context) {
	var cat model.ArticleCategory

	err := c.ShouldBind(&cat)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := c2.DB.Create(&cat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": cat}, "创建成功")

}

func (c2 CategoryController) Update(c *gin.Context) {

	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var cat model.ArticleCategory

	var newCat model.ArticleCategory
	c.ShouldBind(&newCat)

	if err := c2.DB.Where("id=?", catid).First(&cat).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := c2.DB.Model(&cat).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": cat}, "更改成功")
}

func (c2 CategoryController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var cat model.ArticleCategory

	if err := c2.DB.Where("id=?", id).First(&cat).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": cat}, "查询成功")

}

func (c2 CategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	cat := model.ArticleCategory{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	//if err := c2.DB.Debug().Unscoped().Delete(&cat).Error; err != nil {\
	if err := c2.DB.Debug().Delete(&cat).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}

func (c2 CategoryController) ShowAll(c *gin.Context) {

	var cats []model.ArticleCategory

	if err := c2.DB.Model(model.ArticleCategory{}).Find(&cats).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": cats}, "查询成功")

}
