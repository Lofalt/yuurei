package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"strconv"
	"strings"
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

	var messages []model.Message

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	m.DB.Model(model.Message{}).Count(&total)
	if err := m.DB.Preload(clause.Associations).Order("created_at desc").Offset(offset).Limit(pageSize).Find(&messages).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{"data": messages, "total": total}, "success")
}

func (m MessageController) Create(c *gin.Context) {
	var msg model.Message

	err := c.ShouldBind(&msg)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if msg.UserName == "" {
		msg.UserName = "匿名具足虫"
		//msg.Icon = ""
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

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	msg := model.Message{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	err := m.DB.First(&msg).Error
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	//if err := c2.DB.Unscoped().Delete(&msg).Error; err != nil {\
	if err := m.DB.Delete(&msg).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}
	pics := strings.Split(msg.Pics, ",")
	for _, item := range pics {
		path := "." + item
		os.Remove(path)
	}
	response.Success(c, gin.H{}, "删除成功")

}
