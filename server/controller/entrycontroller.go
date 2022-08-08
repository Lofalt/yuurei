package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IEntryController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
	GetEntryById(c *gin.Context)
}

type EntryController struct {
	DB *gorm.DB
}

func (e EntryController) GetEntryById(c *gin.Context) {
	id := c.Params.ByName("id")
	nid, _ := strconv.Atoi(id)
	fmt.Println(nid)
	entry := model.Entry{
		Model: gorm.Model{
			ID:        uint(nid),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	e.DB.Preload("Infos").Preload("EntryCategory").Find(&entry)
	response.Success(c, gin.H{
		"data": entry,
	}, "success")
}
func (e EntryController) Create(c *gin.Context) {
	var entry model.Entry
	err := c.ShouldBind(&entry)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(entry)
	result := e.DB.Create(&entry)
	if result.Error != nil {
		response.Fail(c, gin.H{"错误信息": result.Error}, "出错了")
	} else {
		response.Success(c, gin.H{"data": entry}, "添加成功")
	}
}

func (e EntryController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryController) Show(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryController) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryController) ShowAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewEntryController() IEntryController {
	db := common.GetDb()
	db.AutoMigrate(model.Entry{}, model.EntryInfo{})
	return EntryController{db}
}
