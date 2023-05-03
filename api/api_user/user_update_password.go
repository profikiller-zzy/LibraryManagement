package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/service/user_service"
	"LibraryManagement/utils/jwt_util"
	"LibraryManagement/utils/pwd"
	"github.com/gin-gonic/gin"
	"time"
)

type UpdatePwdRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入新密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入旧密码"`     // 新密码
}

func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var upReq UpdatePwdRequest
	err := c.ShouldBindJSON(&upReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &upReq, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)
	var user model.UserModel
	err = global.Db.First(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	ok := pwd.VerifyPwd(upReq.OldPwd, user.Password)
	if !ok {
		response.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := pwd.BcryptPw(upReq.Pwd)
	err = global.Db.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("修改密码失败", c)
		return
	}
	response.OKWithMessage("修改密码成功，清牢记你的密码", c)
	tokenString := c.Request.Header.Get("token")
	// 计算距离这个token失效还有多长时间
	duration := claims.ExpiresAt - time.Now().Unix()
	err = user_service.UserService{}.AddInvalidTokenToBlackList(tokenString, time.Duration(duration))
	if err != nil {
		global.Log.Error(err.Error())
	}
	return
}
