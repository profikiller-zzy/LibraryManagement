package router

import (
	"LibraryManagement/api/api_admin"
	middleware "LibraryManagement/middleware/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var AdminCaptchaStore = cookie.NewStore([]byte("EKfGVamg0nM6WlHK"))

func (r RGroup) AdminRouter() {
	var AdminApiApp api_admin.AdminApi
	r.Use(sessions.Sessions("captcha", AdminCaptchaStore))
	//r.POST("/admin_update/", middleware.JwtAdmin(), AdminApiApp.UpdateAdminView)
	r.POST("/admin_login/", AdminApiApp.AdminLogin)
	r.POST("/admin_binding_email/", middleware.JwtAdmin(), AdminApiApp.AdminEmailBindingView)
	r.POST("/admin_update_pwd/", middleware.JwtAdmin(), AdminApiApp.UpdateAdminView)
}
