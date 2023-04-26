package admin_service

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/utils/pwd"
)

func CreateAdmin(adminName, password string) error {
	hashPwd := pwd.BcryptPw(password)
	err := global.Db.Create(&model.AdminModel{
		AdminName: adminName,
		Password:  hashPwd,
	}).Error
	return err
}
