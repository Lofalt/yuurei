package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lofalt/yuurei/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	Init()
	r := gin.Default()
	r.StaticFS("/img", http.Dir("./img"))
	// r.Static("/img", "./img")
	common.InitDB()
	r = CollecRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务

}

func Init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("发生错误：%v", err)
		return
	}
}
