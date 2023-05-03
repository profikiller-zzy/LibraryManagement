package api_admin

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"LibraryManagement/utils/pwd"
	"github.com/gin-gonic/gin"
)

type AdminLoginRequest struct {
	AdminName string `json:"admin_name" binding:"required" msg:"请输入管理员用户名"` // 管理员名称
	Password  string `json:"password" binding:"required" msg:"请输入管理员密码"`    // 密码
}

func (AdminApi) AdminLogin(c *gin.Context) {
	var adminLoginReq AdminLoginRequest
	err := c.ShouldBindJSON(&adminLoginReq)
	// 校验参数
	if err != nil {
		response.FailBecauseOfParamError(err, &adminLoginReq, c)
		return
	}

	// 验证管理员是否存在
	var adminModel model.AdminModel
	err = global.Db.Take(&adminModel, "admin_name = ?", adminLoginReq.AdminName).Error
	if err != nil { // 管理员不存在
		global.Log.Warnln("用户名不存在")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码是否正确
	pwdIsCorrect := pwd.VerifyPwd(adminLoginReq.Password, adminModel.Password)
	if !pwdIsCorrect {
		global.Log.Warnln("用户名密码错误")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 验证成功，生成token
	tokenString, err := jwt_util.GenerateTokenForAdmin(jwt_util.JwtAdminPayLoad{
		AdminID: adminModel.ID,
	})
	if err != nil {
		global.Log.Warnln(err.Error())
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	response.OKWithData(tokenString, c)
}
