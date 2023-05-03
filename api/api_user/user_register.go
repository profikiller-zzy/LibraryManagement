package api_user

import (
	"LibraryManagement/model/response"
	"LibraryManagement/service/user_service"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
	NickName string `json:"nick_name"` // 昵称
}

func (UserApi) UserRegisterView(c *gin.Context) {
	var urReq UserRegisterRequest
	if err := c.ShouldBindJSON(&urReq); err != nil {
		response.FailBecauseOfParamError(err, &urReq, c)
		return
	}

	err := user_service.UserService{}.CreateUser(urReq.UserName, urReq.NickName, urReq.Password)
	if err != nil { // 用户创建失败
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("用户注册成功", c)
}
