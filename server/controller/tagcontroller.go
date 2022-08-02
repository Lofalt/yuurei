package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ITagController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type TagController struct {
	DB *gorm.DB
}

func (t TagController) Create(c *gin.Context) {
	var tag model.Tag

	err := c.ShouldBind(&tag)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := t.DB.Create(&tag).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": tag}, "创建成功")
}

func (t TagController) Update(c *gin.Context) {
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var tag model.Tag

	var newCat model.Tag
	c.ShouldBind(&newCat)

	if err := t.DB.Where("id=?", catid).First(&tag).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := t.DB.Model(&tag).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": tag}, "更改成功")
}

func (t TagController) Show(c *gin.Context) {
	id := c.Params.ByName("name")

	//var tag vo.TagRequest
	//var actualTag model.Tag
	tag := model.Tag{
		TagName: id,
	}
	c.ShouldBind(&tag)
	if err := t.DB.Where("tag_name=?", tag.TagName).First(&tag).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": tag}, "查询成功")
}

func (t TagController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	tag := model.Tag{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	//if err := t.DB.Unscoped().Delete(&tag).Error; err != nil {\
	if err := t.DB.Delete(&tag).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}

func NewTagController() ITagController {
	db := common.GetDb()
	db.AutoMigrate(model.Tag{})
	return TagController{db}
}

func (t TagController) ShowAll(c *gin.Context) {
	var cats []model.Tag

	if err := t.DB.Model(model.Tag{}).Find(&cats).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": cats}, "查询成功")

}
