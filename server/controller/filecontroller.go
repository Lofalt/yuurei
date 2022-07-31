package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"log"
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
	timeStame := strconv.Itoa(int(time.Now().UnixMilli())) + "." + strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1]
	filename := "./img/" + timeStame
	fmt.Println(filename)
	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	responFileName := "/img/" + timeStame
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
