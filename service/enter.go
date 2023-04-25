package service

import "LibraryManagement/service/user_service"

type Service struct {
	UserServiceApp user_service.UserService
}

var ServiceApp = new(Service)
