package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type IEntryController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
	GetEntryById(c *gin.Context)
	GetEntriesByCategoryID(c *gin.Context)
}

type EntryController struct {
	DB *gorm.DB
}

func (e EntryController) GetEntriesByCategoryID(c *gin.Context) {
	id := c.Params.ByName("id")
	tid, _ := strconv.Atoi(id)
	entry := model.Entry{
		EntryCategoryID: uint(tid),
		EntryCategory:   model.EntryCategory{},
		Infos:           nil,
	}
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	e.DB.Model(model.Entry{}).Where(&entry).Count(&total)
	var entries []model.Entry
	e.DB.Preload(clause.Associations).Where(&entry).Offset(offset).Order("created_at desc").Limit(pageSize).Find(&entries)
	response.Success(c, gin.H{"data": entries, "total": total}, "success")
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
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var ent model.Entry

	var newEnt model.Entry
	err3 := c.ShouldBind(&newEnt)

	newEnt.ID = uint(catid)
	if err3 != nil {
		response.Fail(c, gin.H{}, err3.Error())
		return
	}

	if err := e.DB.Preload(clause.Associations).Where("id=?", catid).First(&ent).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := e.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&newEnt).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}
	err := e.DB.Model(&ent).Association("Infos").Replace(newEnt.Infos, ent.Infos)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": newEnt}, "更改成功")
}

func (e EntryController) Show(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e EntryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	entry := model.Entry{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//if err := c2.DB.Unscoped().Delete(&cat).Error; err != nil {\
	if err := e.DB.Delete(&entry).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}

func (e EntryController) ShowAll(c *gin.Context) {
	var entries []model.Entry

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	e.DB.Model(model.Entry{}).Count(&total)
	if err := e.DB.Preload(clause.Associations).Order("created_at desc").Offset(offset).Limit(pageSize).Find(&entries).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{"data": entries, "total": total}, "success")
}

func NewEntryController() IEntryController {
	db := common.GetDb()
	db.AutoMigrate(model.Entry{}, model.EntryInfo{})
	return EntryController{db}
}
