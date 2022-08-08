package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IEntryCategoryController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type EntryCategoryController struct {
	DB *gorm.DB
}

func (e EntryCategoryController) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryCategoryController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryCategoryController) Show(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryCategoryController) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryCategoryController) ShowAll(c *gin.Context) {
	var cats []model.EntryCategory

	if err := e.DB.Model(model.EntryCategory{}).Find(&cats).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": cats}, "查询成功")
}

func NewEntryCategoryController() IEntryCategoryController {
	db := common.GetDb()
	db.AutoMigrate(model.EntryCategory{})
	return EntryCategoryController{db}
}
