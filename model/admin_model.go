package model

type AdminModel struct {
	MODEL
	AdminName string `json:"admin_name"` // 管理员名称
	Password  string `json:"password"`   // 密码
	Email     string `json:"email"`      // 邮箱地址
}
