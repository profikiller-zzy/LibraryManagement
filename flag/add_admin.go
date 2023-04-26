package flag

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/service/admin_service"
	"fmt"
)

func CreateAdmin() {
	var (
		adminName string
		password  string
	)
	fmt.Printf("请输入管理员用户名:")
	fmt.Scan(&adminName)
	// 查询该管理员用户名是否存在
	var admin model.AdminModel
	err := global.Db.First(&admin, "admin_name = ?", adminName).Error
	if err == nil {
		fmt.Printf("该用户名已存在")
		return
	}
	fmt.Printf("请输入管理员密码:")
	fmt.Scan(&password)
	err = admin_service.CreateAdmin(adminName, password)
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Info(fmt.Sprintf("创建管理员%s成功", adminName))
}
