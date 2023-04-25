package jwt_util

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	NickName string `json:"nick_name"`
}

var JwtSecretKey []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
