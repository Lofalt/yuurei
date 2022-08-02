package dto

import "github.com/Lofalt/yuurei/model"

type UserDto struct {
	Name     string `json:"Name"`
	Username string `json:"Username"`
	IsAdmin  bool   `json:"IsAdmin"`
}

func GetUserDto(user model.User) UserDto {
	return UserDto{
		Name:     user.Name,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	}
}
