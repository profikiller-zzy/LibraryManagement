package user_service

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/utils/pwd"
	"github.com/pkg/errors"
)

func (UserService) CreateUser(userName, nickName, password string) error {
	// 判断用户名是否存在
	var userModel model.UserModel
	err := global.Db.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashPwd := pwd.BcryptPw(password)

	// 创建用户
	err = global.Db.Create(&model.UserModel{
		NickName: nickName,
		UserName: userName,
		Password: hashPwd,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
