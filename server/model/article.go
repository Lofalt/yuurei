package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleTitle      string `gorm:"NOT NULL;type:varchar(30)"`
	ArticleContent    string `gorm:"NOT NULL"`
	ArticleReadTimes  int
	ArticleLoveTimes  int
	IsActive          bool             `gorm:"DEFAULT:1"`
	ArticleSummary    string           `gorm:"NOT NULL;type:varchar(50)"`
	ArticleCategoryID int              `gorm:"NOT NULL"`
	ArticleCategory   *ArticleCategory `gorm:"Not null"`
	Tags              []*Tag           `gorm:"many2many:tag_article_tables"`
	ArticleComments   []ArticleComment
}
type ArticleCategory struct {
	gorm.Model
	ArticleCategoryName string `gorm:"type:varchar(20)"`
}

type Tag struct {
	gorm.Model
	TagName  string     `gorm:"type:varchar(30)"`
	TagColor string     `gorm:"type:varchar(30)"`
	Articles []*Article `gorm:"many2many:tag_article_tables"`
}

type ArticleComment struct {
	gorm.Model
	ArticleCommentContent  string `gorm:"NOT NULL"`
	ArticleCommentUserName string `gorm:"NOT NULL;type:varchar(30)"`
	ArticleID              uint
}
