package dto

import "github.com/Lofalt/yuurei/model"

type UserDto struct {
	Username  string `json:"name"`
	Telephone string `json:"telephone"`
}

func GetUserDto(user model.User) UserDto {
	return UserDto{
		Username:  user.Name,
		Telephone: user.Telephone,
	}
}
