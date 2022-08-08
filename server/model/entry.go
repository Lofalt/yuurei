package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Title           string `gorm:"type:varchar(20);not null"`
	Quote           string `gorm:"type:varchar(100);not null"`
	Summary         string `gorm:"type:varchar(100);not null"`
	Detail          string `gorm:"not null"`
	MyWord          string `gorm:"not null"`
	About           string `gorm:"not null"`
	HeaderPicture   string `gorm:"type:varchar(100)"`
	Infos           []EntryInfo
	MainColor       string `gorm:"type:varchar(15)"`
	SecColor        string `gorm:"type:varchar(15)"`
	ThirdColor      string `gorm:"type:varchar(15)"`
	EntryCategoryID uint
	EntryCategory   EntryCategory
}

type EntryInfo struct {
	gorm.Model
	EntryId uint
	Label   string `gorm:"type:varchar(50)"`
	Value   string `gorm:"type:varchar(50)"`
}

type EntryCategory struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20)"`
}
