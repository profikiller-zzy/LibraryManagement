package router

import (
	"LibraryManagement/api/api_user"
	middleware "LibraryManagement/middleware/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var UserCaptchaStore = cookie.NewStore([]byte("EKfGVamg0nM6WlHK"))

func (r RGroup) UserRouter() {
	var UserApiApp api_user.UserApi
	r.Use(sessions.Sessions("captcha", UserCaptchaStore))
	r.POST("/user_login/", UserApiApp.UserLoginView)
	r.POST("/user_register/", UserApiApp.UserRegisterView)
	r.PUT("/user_update/", middleware.JwtAuth(), UserApiApp.UserUpdateView)
	r.PUT("/user_update_password/", middleware.JwtAuth(), UserApiApp.UserUpdatePasswordView)
	r.GET("/user_borrow_record_list/", middleware.JwtAuth(), UserApiApp.BorrowRecordList)
	r.GET("/user_info/", middleware.JwtAuth(), UserApiApp.UserInfoView)
	r.POST("/user_logout/", middleware.JwtAuth(), UserApiApp.UserLogoutView)
	r.GET("/user_list/", middleware.JwtAdmin(), UserApiApp.UserListView)
	r.DELETE("/user_delete/:id", middleware.JwtAdmin(), UserApiApp.DeleteUserView)
}
