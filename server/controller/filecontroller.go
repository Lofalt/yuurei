package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/response"
	"github.com/Lofalt/yuurei/util"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func UploadImg(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	quality, _ := strconv.Atoi(c.Query("qua"))
	dir := c.Query("dir")

	fmt.Println(quality)
	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)
	suffix := strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]

	timeStamp := strconv.Itoa(int(time.Now().UnixMilli())) + "." + suffix

	filepath := "./img/" + dir + "/"
	if exit, _ := util.HasDir(filepath); exit == false {
		os.Mkdir(filepath, os.ModePerm)
	}
	filename := filepath + timeStamp

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	if quality != 0 && (suffix == "jpg" || suffix == "jpeg") {
		util.Compress(filename, quality)

	}

	responFileName := "/img/" + dir + "/" + timeStamp
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func GalleryPic(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	quality, _ := strconv.Atoi(c.Query("qua"))

	suffix := strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]

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
	suffix = strings.ToUpper(suffix)
	if quality != 0 && (suffix == "JPG" || suffix == "JPEG") {
		util.Compress(filename, quality)

	}

	responFileName := filename
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func CropImg(c *gin.Context) {
	path := c.PostForm("imgPath")
	pathtmp := "." + path
	x, _ := strconv.ParseFloat(c.PostForm("x"), 64)
	y, _ := strconv.ParseFloat(c.PostForm("y"), 64)
	x1, _ := strconv.ParseFloat(c.PostForm("width"), 64)
	y1, _ := strconv.ParseFloat(c.PostForm("height"), 64)

	x0 := int(math.Round(x))
	y0 := int(math.Round(y))
	x01 := int(math.Round(x1))
	y01 := int(math.Round(y1))
	x01 = x0 + x01
	y01 = y0 + y01

	err := util.CropPic(pathtmp, x0, y0, x01, y01)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
	}
	response.Success(c, gin.H{"data": path}, "success")
}
