package jwt_util

import (
	"LibraryManagement/global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// GenerateTokenForUser 根据用户ID和用户昵称为用户产生token
func GenerateTokenForUser(payLoad JwtUserPayLoad) (string, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	// Token的有效时间，默认为两个小时
	expireTime := time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.ExpireTime))
	Claim := CustomUserClaims{
		JwtUserPayLoad: payLoad,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// Issuer表示Token的签发者
			Issuer: "图书馆借阅管理系统",
		},
	}
	// NewWithClaims根据Claims结构体创建Token示例
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	// SignedString 方法根据传入的空接口类型参数key，返回完整的签名令牌。
	reqToken, err := reqClaim.SignedString(JwtSecretKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("jwt token生成失败，错误信息：%s", err.Error()))
	}
	return reqToken, nil
}

// GenerateTokenForAdmin 根据管理员ID和为管理员产生token
func GenerateTokenForAdmin(payLoad JwtAdminPayLoad) (string, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	// Token的有效时间，默认为两个小时
	expireTime := time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.ExpireTime))
	Claim := CustomAdminClaims{
		JwtAdminPayLoad: payLoad,
		StandardClaims: jwt.StandardClaims{
			// 生效时间
			IssuedAt: time.Now().Unix(),
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// Issuer表示Token的签发者
			Issuer: "图书馆借阅管理系统",
		},
	}
	global.Log.Info(expireTime)
	// NewWithClaims根据Claims结构体创建Token示例
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	// SignedString 方法根据传入的空接口类型参数key，返回完整的签名令牌。
	reqToken, err := reqClaim.SignedString(JwtSecretKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("jwt token生成失败，错误信息：%s", err.Error()))
	}
	return reqToken, nil
}

// VerifyTokenForUser 为用户解析和验证token
func VerifyTokenForUser(tokenString string) (*CustomUserClaims, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &CustomUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		global.Log.Error(fmt.Sprintf("Verify token error : %s", err.Error()))
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomUserClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("Invalid token string")
}

// VerifyTokenForAdmin 为管理员解析和验证token
func VerifyTokenForAdmin(tokenString string) (*CustomAdminClaims, error) {
	JwtSecretKey = []byte(global.Config.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &CustomAdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		global.Log.Error(fmt.Sprintf("Verify token error : %s", err.Error()))
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomAdminClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("Invalid token string")
}
