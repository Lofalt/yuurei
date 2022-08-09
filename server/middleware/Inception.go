package middleware

import (
	"fmt"
	"github.com/Lofalt/yuurei/util"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func Inception() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//获取header
		//验证格式
		ip := ctx.ClientIP()
		//path := ctx.FullPath()
		conn := util.GetRedisConn()
		defer conn.Close()
		str, _ := redis.String(conn.Do("Get", "bill"))
		conn.Do("Set", "bill", time.Now().UnixMilli())
		conn.Do("expire", "bill", 5)
		fmt.Println(ip + str)
		//ctx.Set("user", user)

		ctx.Next()
	}
}
