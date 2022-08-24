package middleware

import (
	"fmt"
	"github.com/Lofalt/yuurei/response"
	"strings"

	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//获取header
		tokenString := ctx.GetHeader("Authorization")
		//验证格式
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Fail(ctx, gin.H{}, "权限不足捏")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			response.Fail(ctx, gin.H{}, "token无效")
			ctx.Abort()
			return

		}
		fmt.Println(claims)
		userId := claims.UserId
		var user model.User
		db := common.GetDb()
		db.Debug().First(&user, "id = ?", userId)
		if user.ID == "" {
			response.Fail(ctx, gin.H{}, "用户不存在")
			ctx.Abort()
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
