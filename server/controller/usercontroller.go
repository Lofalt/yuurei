package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/dto"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/Lofalt/yuurei/util"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Register(c *gin.Context) {

	db := common.GetDb()
	conn := util.GetRedisConn()
	defer conn.Close()
	name := c.PostForm("name")
	password := c.PostForm("password")
	username := c.PostForm("username")
	validCode := c.PostForm("code")
	code, err := redis.String(conn.Do("Get", username))

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
	if code == "" {
		response.Fail(c, gin.H{}, "请先发送验证码")
		return
	}
	if validCode != code {
		response.Fail(c, gin.H{}, "验证码错误")
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

func SendEmail(c *gin.Context) {
	db := common.GetDb()
	username := c.PostForm("Email")
	var user model.User
	db.Where("UserName=?", username).First(&user)
	if user.ID != "" {
		response.Fail(c, gin.H{}, "邮箱已被注册")
		return
	}
	code, err := SendEmailValidate([]string{username})
	if err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}
	conn := util.GetRedisConn()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	conn.Do("Set", username, code)
	conn.Do("expire", username, 300)
	response.Success(c, gin.H{}, "waiting")

}

func SendEmailValidate(em []string) (string, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("Lofalt <515636512@qq.com>")
	e.To = em
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`
<html>
	<div>
		<div>
			你好
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>你于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
</html>
	`, t, vCode)
	e.HTML = []byte(content)
	e.Subject = "接收验证码"
	e.AttachFile("../img/MyIcon/2cd5dfc2-28bf-4d6e-944b-9c605f44ee321659793639194.jpg")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "515636512@qq.com", "knjdcbozvueibhhf", "smtp.qq.com"))
	return vCode, err
}

func Report(ctx *gin.Context) {
	//获取header
	tokenString := ctx.GetHeader("Authorization")
	var userId string
	IPAddressLoc := ""
	if tokenString != "" {
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {

		}
		userId = claims.UserId
		var user model.User
		db := common.GetDb()
		db.Debug().First(&user, "id = ?", userId)

	}
	post, err := http.Post("https://apikey.net/?type=json&ip="+ctx.ClientIP(), "application/json", bytes.NewReader([]byte{}))
	if err != nil {
		log.Println(err.Error())
		return
	}
	if post.Status == "200 OK" {

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(post.Body)
		body, _ := ioutil.ReadAll(post.Body)
		if err != nil {
			return
		}
		var res IPResult
		err = json.Unmarshal(body, &res)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		IPAddressLoc = res.Address
	}
	visitor := model.Visitor{
		Model:        gorm.Model{},
		UserID:       userId,
		IPAddress:    ctx.ClientIP(),
		IPAddressLoc: IPAddressLoc,
	}

	db := common.GetDb()
	var vis model.Visitor
	if userId != "" {
		db.Debug().Where(" user_id = ?", userId).Preload("User").First(&vis)
		if vis.ID != 0 {
			db.Where("id = ?", vis.ID).Updates(&visitor)
			response.Success(ctx, gin.H{"data": vis}, "success")
			return
		}

	}
	db.Debug().Where(" ip_address = ?", ctx.ClientIP()).Preload("User").First(&vis)
	if vis.ID != 0 {
		db.Where("id = ?", vis.ID).Updates(&visitor)
		response.Success(ctx, gin.H{"data": vis}, "success")
		return
	} else {
		db.Create(&visitor)
	}
}

func ShowVisitors(c *gin.Context) {
	var visitors []model.Visitor
	db := common.GetDb()

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var total int64
	db.Model(model.Visitor{}).Count(&total)
	if err := db.Preload(clause.Associations).Order("updated_at desc").Offset(offset).Limit(pageSize).Find(&visitors).Error; err != nil {
		response.Fail(c, gin.H{}, err.Error())
		return
	}

	response.Success(c, gin.H{"data": visitors, "total": total}, "success")
}

type IPResult struct {
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Ip      string `json:"ip"`
	Iplong  int    `json:"iplong"`
	Address string `json:"address"`
	Ipproxy bool   `json:"ipproxy"`
}
