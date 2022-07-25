package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"size:255"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
}

// func NewUser(name, password, telephone string) *User {
// 	return &User{
// 		name,
// 		password,
// 		telephone,
// 	}
// }
