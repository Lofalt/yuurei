package controller

import (
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"log"
)

func UploadImgController(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)
	err := c.SaveUploadedFile(file, "./img/"+file.Filename)
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	responFileName := "/img/" + file.Filename
	response.Success(c, gin.H{
		"fileName": responFileName,
	}, "success")
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
