package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleTitle         string `gorm:"NOT NULL"`
	ArticleContent       string `gorm:"NOT NULL"`
	ArticleReadTimes     int
	ArticleLoveTimes     int
	IsActive             bool            `gorm:"DEFAULT:1"`
	ArticleSummary       string          `gorm:"NOT NULL"`
	ArticleCategoryRefer int             `gorm:"NOT NULL"`
	ArticleCategory      ArticleCategory `gorm:"foreignKey:ArticleCategoryRefer"`
	Tags                 []*Tag          `gorm:"many2many:tag_article_tables"`
	ArticleComments      []ArticleComment
}
type ArticleCategory struct {
	gorm.Model
	ArticleCategoryName string
}

type Tag struct {
	gorm.Model
	TagName  string
	TagColor string
	Articles []*Article `gorm:"many2many:tag_article_tables"`
}

type ArticleComment struct {
	gorm.Model
	ArticleCommentContent  string `gorm:"NOT NULL"`
	ArticleCommentUserName string `gorm:"NOT NULL"`
	ArticleID              uint
}
