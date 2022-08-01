package model

import "gorm.io/gorm"

type PictureEntry struct {
	gorm.Model
	Description string `gorm:"type:varchar(50);not null"`
	Path        string `gorm:"type:varchar(100);not null"`
}
