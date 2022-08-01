package controller

import (
	"github.com/Lofalt/yuurei/response"
	"github.com/Lofalt/yuurei/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func UploadImgController(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)
	timeStamp := strconv.Itoa(int(time.Now().UnixMilli())) + "." + strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]

	filepath := "./img/"
	if exit, _ := util.HasDir(filepath); exit == false {
		os.Mkdir(filepath, os.ModePerm)
	}
	filename := filepath + timeStamp

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	responFileName := "/img/" + timeStamp
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func GalleryPic(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)
	timeStamp := strconv.Itoa(int(time.Now().UnixMilli())) + "." + strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]

	filepath := "./img/gallery"
	if exit, _ := util.HasDir(filepath); exit == false {
		os.Mkdir(filepath, os.ModePerm)
	}
	filename := filepath + "/" + timeStamp

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	responFileName := filename
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
