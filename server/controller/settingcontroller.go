package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ISettingController interface {
	RESTfulcontroller
}

type SettingController struct {
	DB *gorm.DB
}

func NewSettingController() ISettingController {
	db := common.GetDb()
	db.AutoMigrate(model.GlobalConfig{})
	return SettingController{DB: db}
}

func (s SettingController) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s SettingController) Update(c *gin.Context) {
	setid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var setting model.GlobalConfig

	var newset model.GlobalConfig
	c.ShouldBind(&newset)

	if err := s.DB.Where("id=?", setid).First(&setting).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := s.DB.Model(&setting).Updates(&newset).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": setting}, "更改成功")
}

func (s SettingController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var set model.GlobalConfig

	if err := s.DB.Where("id=?", id).First(&set).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": set}, "查询成功")

}

func (s SettingController) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
