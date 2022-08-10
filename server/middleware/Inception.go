package middleware

import (
	"fmt"
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
	"github.com/Lofalt/yuurei/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func Inception() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//获取header
		//验证格式
		//ip := ctx.ClientIP()
		//path := ctx.FullPath()
		//conn := util.GetRedisConn()
		//获取header
		tokenString := ctx.GetHeader("Authorization")
		//验证格式
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(500, gin.H{"message": "权限不足捏"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(500, gin.H{"message": "token无效"})
			ctx.Abort()
			return

		}
		fmt.Println(claims)
		userId := claims.UserId
		var user model.User
		db := common.GetDb()
		db.Debug().First(&user, "id = ?", userId)
		if !user.IsAdmin {
			response.Fail(ctx, gin.H{}, "你没有权限")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
