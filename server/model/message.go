package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MessageContent string `gorm:"type:text;not null"`
	UserName       string `gorm:"type:varchar(50);"`
	Icon           string `gorm:"type:varchar(100)"`
}
