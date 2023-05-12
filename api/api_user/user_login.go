package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
	"LibraryManagement/model/response"
	"LibraryManagement/utils/jwt_util"
	"LibraryManagement/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string `json:"password" binding:"required" msg:"请输入密码"`   // 密码
}

// UserLoginView 用户注册接口
func (UserApi) UserLoginView(c *gin.Context) {
	var ulReq UserLoginRequest
	err := c.ShouldBindJSON(&ulReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	// 验证用户是否存在
	var userModel model.UserModel
	err = global.Db.Take(&userModel, "user_name = ?", ulReq.UserName).Error
	if err != nil { // 该用户不存在
		global.Log.Warnln("用户名不存在")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码是否正确
	pwdIsCorrect := pwd.VerifyPwd(ulReq.Password, userModel.Password)
	if !pwdIsCorrect {
		global.Log.Warnln("用户名密码错误")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 验证成功，生成Token
	tokenString, err := jwt_util.GenerateTokenForUser(jwt_util.JwtUserPayLoad{
		UserID:   userModel.ID,
		NickName: userModel.NickName,
	})
	if err != nil {
		global.Log.Warnln(err.Error())
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	response.OKWithData(tokenString, c)
}
