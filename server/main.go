package main

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	Init()
	r := gin.Default()
	r.StaticFS("/img", http.Dir("./img"))
	// r.Static("/img", "./img")
	common.InitDB()
	r = CollecRouter(r)
	port := viper.GetString("server.port")
	//db := common.GetDb()

	//tag1 := model.Tag{
	//	TagName: "fuck",
	//}
	//tag3 := model.Tag{
	//	TagName: "man",
	//}
	//tag4 := model.Tag{TagName: "haha"}
	//tag2 := model.Tag{TagName: "you"}
	//c1 := model.ArticleCategory{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	ArticleCategoryName: "玩玩玩就知道玩",
	//}
	//c2 := model.ArticleCategory{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	ArticleCategoryName: "没半点出息",
	//}
	//a1 := model.Article{
	//	ID:                   0,
	//	CreatedAt:            time.Time{},
	//	UpdatedAt:            time.Time{},
	//	DeletedAt:            gorm.DeletedAt{},
	//	ArticleTitle:         "",
	//	ArticleContent:       "",
	//	ArticleReadTimes:     0,
	//	ArticleLoveTimes:     0,
	//	IsActive:             false,
	//	ArticleSummary:       "",
	//	ArticleCategoryRefer: 1,
	//	Tags:                 []*model.Tag{&tag1, &tag2, &tag3, &tag4},
	//}
	//err := db.AutoMigrate(&model.Article{}, &model.ArticleCategory{}, &model.Tag{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//db.Create(&c1)
	//db.Create(&c2)
	//db.Debug().Create(&a1)
	//fmt.Println(db)

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
