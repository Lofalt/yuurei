package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50)"`
	Link    string `gorm:"type:varchar(100)"`
	Icon    string `gorm:"type:varchar(200)"`
	Summary string `gorm:"type:varchar(100)"`
	Color   string `gorm:"type:varchar(10)"`
}
