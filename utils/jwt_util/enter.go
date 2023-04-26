package jwt_util

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtUserPayLoad struct {
	UserID   uint   `json:"user_id"`
	NickName string `json:"nick_name"`
}

type JwtAdminPayLoad struct {
	AdminID uint `json:"admin_id"`
}

var JwtSecretKey []byte

type CustomUserClaims struct {
	JwtUserPayLoad
	jwt.StandardClaims
}

type CustomAdminClaims struct {
	JwtAdminPayLoad
	jwt.StandardClaims
}
