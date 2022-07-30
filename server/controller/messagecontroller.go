package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type IMessageController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type MessageController struct {
	DB *gorm.DB
}

func NewMessageController() IMessageController {
	db := common.GetDb()
	db.AutoMigrate(model.Message{})
	return MessageController{db}
}

func (m MessageController) ShowAll(c *gin.Context) {

	var msgs []model.Message

	if err := m.DB.Model(model.Message{}).Order("created_at desc").Find(&msgs).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": msgs}, "查询成功")
}

func (m MessageController) Create(c *gin.Context) {
	var msg model.Message

	err := c.ShouldBind(&msg)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := m.DB.Create(&msg).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": msg}, "创建成功")
}

func (m MessageController) Update(c *gin.Context) {

	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var msg model.Message

	var newCat model.Message
	c.ShouldBind(&newCat)

	if err := m.DB.Where("id=?", catid).First(&msg).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := m.DB.Model(&msg).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": msg}, "更改成功")
}

func (m MessageController) Show(c *gin.Context) {

	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var msg model.Message

	var newCat model.Message
	c.ShouldBind(&newCat)

	if err := m.DB.Where("id=?", catid).First(&msg).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := m.DB.Model(&msg).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": msg}, "更改成功")
}

func (m MessageController) Delete(c *gin.Context) {

	var cats []model.Message

	if err := m.DB.Model(model.Message{}).Find(&cats).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": cats}, "查询成功")

}
