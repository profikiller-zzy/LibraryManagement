package router

import (
	"LibraryManagement/api/api_admin"
	middleware "LibraryManagement/middleware/jwt"
)

func (r RGroup) AdminRouter() {
	var AdminApiApp api_admin.AdminApi
	r.POST("/admin_update/", middleware.JwtAdmin(), AdminApiApp.UpdateAdminView)
	r.POST("/admin_login/", AdminApiApp.AdminLogin)
}
