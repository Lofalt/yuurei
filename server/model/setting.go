package model

import "gorm.io/gorm"

type GlobalConfig struct {
	gorm.Model
	WebsiteName           string `gorm:"type:varchar(20)"`
	MyName                string `gorm:"type:varchar(20)"`
	MessageDefaultIcon    string `gorm:"type:varchar(100)"`
	AnonymousName         string `gorm:"type:varchar(100)"`
	Icon                  string `gorm:"type:varchar(100)"`
	BackgroundImage       string `gorm:"type:varchar(100)"`
	MobileBackgroundImage string `gorm:"type:varchar(100)"`
	DateColor             string `gorm:"type:varchar(10)"`
	ButtonColor           string `gorm:"type:varchar(10)"`
	SecColor              string `gorm:"type:varchar(10)"`
	ThirdColor            string `gorm:"type:varchar(10)"`
	FourthColor           string `gorm:"type:varchar(10)"`
	Slogan                string `gorm:"type:varchar(200)"`
}
