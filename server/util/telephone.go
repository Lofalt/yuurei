package util

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
)

func UserNameExist(username string) bool {
	db := common.GetDb()
	var user *model.User
	db.Debug().First(&user, "username = ?", username)
	return user.ID != ""
}
