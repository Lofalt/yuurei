package controller

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/dto"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/Lofalt/yuurei/util"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Register(c *gin.Context) {

	db := common.GetDb()
	name := c.PostForm("name")
	password := c.PostForm("password")
	username := c.PostForm("username")
	errR := db.AutoMigrate(&model.User{})
	if errR != nil {
		return
	}
	if name == "" {
		name = util.RandomString(10)
	}

	if util.UserNameExist(username) {
		response.Response(c, 200, 400, nil, "当前用户名已被注册")
		return
	}
	//加密密码
	passwordString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, 500, -1, nil, "服务器端错误")

	}
	db.Create(&model.User{
		ID:   uuid.NewV4().String(),
		Name: name, Password: string(passwordString), Username: username})
	response.Success(c, nil, "恭喜注册成功")
}

func Login(c *gin.Context) {
	db := common.GetDb()
	password := c.PostForm("password")
	username := c.PostForm("username")
	//验证手机号是否存在
	if !util.UserNameExist(username) {
		response.Response(c, 200, 400, nil, "找不到账号")

		return
	}
	var user *model.User
	db.First(&user, "username = ?", username)

	//判断密码是否错误
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密码错误")
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(c, 500, -1, nil, "系统错误")
		log.Printf("token generate error:%v", err)
		return
	}
	response.Success(c, gin.H{"token": token, "data": user}, "登陆成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"user": dto.GetUserDto(user.(model.User))}})

}

func Imgs(ctx *gin.Context) {
	picList := make([]string, 0)

	dir := "img"

	absolutpath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(absolutpath)
	namelist, _ := os.ReadDir(dir)

	for _, name := range namelist {
		tmpname := "/" + dir + "/" + name.Name()
		// fmt.Println(tmpname)
		picList = append(picList, tmpname)
		// fmt.Println(picList)

	}

	picnum, _ := strconv.Atoi(ctx.PostForm("num"))

	response.Success(ctx, gin.H{"list": picList[0:picnum]}, "ok")

}
