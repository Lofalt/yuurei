package model

import "gorm.io/gorm"

type GlobalConfig struct {
	gorm.Model
	WebsiteName           string `gorm:"type:varchar(20)"`
	MyName                string `gorm:"type:varchar(20)"`
	MessageDefaultIcon    string `gorm:"type:varchar(100)"`
	Icon                  string `gorm:"type:varchar(100)"`
	BackgroundImage       string `gorm:"type:varchar(100)"`
	MobileBackgroundImage string `gorm:"type:varchar(100)"`
	DateColor             string `gorm:"type:varchar(10)"`
	ButtonColor           string `gorm:"type:varchar(10)"`
}
