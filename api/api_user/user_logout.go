package api_user

import (
	"LibraryManagement/global"
	"LibraryManagement/model/response"
	"LibraryManagement/service"
	"LibraryManagement/utils/jwt_util"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt_util.CustomUserClaims)

	tokenString := c.Request.Header.Get("token")

	// 需要现在距离过期时间还有多久，以这个计算出的时间来当作该记录的有效时间
	expiresAt := time.Unix(claims.ExpiresAt, 0)
	duration := expiresAt.Sub(time.Now())
	err := service.ServiceApp.UserServiceApp.AddInvalidTokenToBlackList(tokenString, duration)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("注销失败", c)
		return
	}
	response.OKWithMessage("注销成功", c)
}
