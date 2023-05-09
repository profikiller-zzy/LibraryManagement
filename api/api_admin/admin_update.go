package api_admin

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"LibraryManagement/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePwdRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

// UpdateAdminPwdView 管理员修改密码
func (AdminApi) UpdateAdminPwdView(c *gin.Context) {
	var upReq UpdatePwdRequest
	err := c.ShouldBindJSON(&upReq)
	// 判断跳转链接是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &upReq, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomAdminClaims)
	var adminModel model.AdminModel
	err = global.Db.First(&adminModel, claims.AdminID).Error
	if err != nil {
		response.FailWithMessage("管理员不存在", c)
		return
	}
	ok := pwd.VerifyPwd(upReq.OldPwd, adminModel.Password)
	if !ok {
		response.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := pwd.BcryptPw(upReq.Pwd)
	err = global.Db.Model(&adminModel).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("修改密码失败", c)
		return
	}
	response.OKWithMessage("修改密码成功，清牢记你的密码", c)
	return
}
