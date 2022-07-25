package model

import "time"
import "gorm.io/gorm"

type Article struct {
	ID               uint           `gorm:"column:article_id"`
	CreatedAt        time.Time      `gorm:"column:article_create_time"`
	UpdatedAt        time.Time      `gorm:"column:article_last_mod_time"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	ArticleTitle     string
	ArticleContent   string
	ArticleReadTimes int
	ArticleLoveTimes int
	IsActive         bool
	ArticleSummary   string
}
type Article2 struct {
	ID               uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	ArticleTitle     string
	ArticleContent   string
	ArticleReadTimes int
	ArticleLoveTimes int
	IsActive         bool
	ArticleSummary   string
}
