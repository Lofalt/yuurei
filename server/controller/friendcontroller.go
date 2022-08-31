package controller

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type IFriendController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type FriendController struct {
	DB *gorm.DB
}

func (f FriendController) Create(c *gin.Context) {
	var friend model.Friend

	err := c.ShouldBind(&friend)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := f.DB.Create(&friend).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": friend}, "创建成功")
}

func (f FriendController) Update(c *gin.Context) {
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var friend model.Friend

	var newFriend model.Friend
	c.ShouldBind(&newFriend)

	if err := f.DB.Where("id=?", catid).First(&friend).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := f.DB.Model(&friend).Updates(&newFriend).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": friend}, "更改成功")
}

func (f FriendController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var friend model.Friend

	if err := f.DB.Where("id=?", id).First(&friend).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": friend}, "查询成功")
}

func (f FriendController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	friend := model.Friend{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	//if err := f.DB.Unscoped().Delete(&friend).Error; err != nil {\
	if err := f.DB.Where("id=?", id).First(&friend).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err := f.DB.Delete(&friend).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}
	err := os.Remove("." + friend.Icon)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())

		return
	}
	response.Success(c, gin.H{}, "删除成功")
}

func (f FriendController) ShowAll(c *gin.Context) {
	var friends []model.Friend

	if err := f.DB.Model(model.Friend{}).Find(&friends).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": friends}, "查询成功")
}

func NewFriendController() IFriendController {
	db := common.GetDb()
	db.AutoMigrate(model.Friend{})
	return FriendController{db}
}
