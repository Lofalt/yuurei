package util

import (
	"github.com/Lofalt/yuurei/common"
	"github.com/Lofalt/yuurei/model"
)

func TelephoneExit(telephone string) bool {
	db := common.GetDb()
	var user *model.User
	db.First(&user, "telephone = ?", telephone)
	return user.ID != 0
}
