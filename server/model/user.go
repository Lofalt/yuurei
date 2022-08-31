package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"type:varchar(50);not null"`
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"size:255"`
	Username string `gorm:"type:varchar(30);not null;unique"`
	IsAdmin  bool   `gorm:"DEFAULT:0"`
}

// func NewUser(name, password, telephone string) *User {
// 	return &User{
// 		name,
// 		password,
// 		telephone,
// 	}
// }
type Visitor struct {
	gorm.Model
	UserID       string `gorm:"type:varchar(50);unique;default:null"`
	Username     string `gorm:"type:varchar(100)"`
	User         User
	IPAddress    string `gorm:"type:varchar(50)"`
	IPAddressLoc string `gorm:"type:varchar(50)"`
}
