package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type IPicController interface {
	RESTfulcontroller
	ShowAll(c *gin.Context)
}

type PicController struct {
	DB *gorm.DB
}

func NewIPicController() IPicController {
	db := common.GetDb()
	db.AutoMigrate(&model.PictureEntry{})
	return PicController{DB: db}
}

func (p PicController) ShowAll(c *gin.Context) {
	var pics []model.PictureEntry
	var total int64
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}

	if err := p.DB.Model(model.PictureEntry{}).Order("created_at desc").Offset(offset).Limit(pageSize).Find(&pics).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	p.DB.Model(model.PictureEntry{}).Count(&total)
	for i := 0; i < len(pics); i++ {
		pics[i].Path = strings.TrimPrefix(pics[i].Path, ".")
	}

	response.Success(c, gin.H{"data": pics, "total": total}, "查询成功")

}

func (p PicController) Create(c *gin.Context) {
	var pic model.PictureEntry

	err := c.ShouldBind(&pic)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if err2 := p.DB.Create(&pic).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": pic}, "创建成功")

}

func (p PicController) Update(c *gin.Context) {
	catid, _ := strconv.Atoi(c.Params.ByName("id"))

	//查找是否存在
	var pic model.PictureEntry

	var newCat model.PictureEntry
	c.ShouldBind(&newCat)

	if err := p.DB.Where("id=?", catid).First(&pic).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	if err2 := p.DB.Model(&pic).Updates(&newCat).Error; err2 != nil {
		response.Fail(c, gin.H{}, err2.Error())
		return
	}

	response.Success(c, gin.H{"data": pic}, "更改成功")
}

func (p PicController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var pic model.PictureEntry

	if err := p.DB.Where("id=?", id).First(&pic).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{"data": pic}, "查询成功")

}

func (p PicController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	pic := model.PictureEntry{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	p.DB.First(&pic)
	fmt.Println(pic)
	os.Remove(pic.Path)
	//if err := p.DB.Unscoped().Delete(&pic).Error; err != nil {\
	if err := p.DB.Delete(&pic).Error; err != nil {

		response.Fail(c, gin.H{}, err.Error())
		return
	}
	response.Success(c, gin.H{}, "删除成功")
}
