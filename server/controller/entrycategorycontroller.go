package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type IEntryCategoryController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type EntryCategoryController struct {
	DB *gorm.DB
}

func (e EntryCategoryController) Create(c *gin.Context) {
	var entryCat model.EntryCategory

	err := c.ShouldBind(&entryCat)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := e.DB.Create(&entryCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": entryCat}, "创建成功")
}

func (e EntryCategoryController) Update(c *gin.Context) {
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var entryCat model.EntryCategory

	var newCat model.EntryCategory
	c.ShouldBind(&newCat)

	if err := e.DB.Where("id=?", catid).First(&entryCat).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := e.DB.Model(&entryCat).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": entryCat}, "更改成功")
}

func (e EntryCategoryController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var entryCat model.EntryCategory

	if err := e.DB.Where("id=?", id).First(&entryCat).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": entryCat}, "查询成功")

}

func (e EntryCategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	entryCat := model.EntryCategory{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	//if err := e.DB.Unscoped().Delete(&entryCat).Error; err != nil {\
	if err := e.DB.Delete(&entryCat).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}
}
func (e EntryCategoryController) ShowAll(c *gin.Context) {
	var entries []model.EntryCategory

	if err := e.DB.Model(model.EntryCategory{}).Find(&entries).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": entries}, "查询成功")
}

func NewEntryCategoryController() IEntryCategoryController {
	db := common.GetDb()
	db.AutoMigrate(model.EntryCategory{})
	return EntryCategoryController{db}
}
