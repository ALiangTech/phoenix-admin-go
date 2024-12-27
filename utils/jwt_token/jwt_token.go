package jwtToken

import (
	"phoenix-go-admin/utils/mistakes"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateJwt(secret string, uid string) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		uid,
		jwt.RegisteredClaims{
			Issuer: "phoenix",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwts, err := token.SignedString([]byte(secret))
	if err != nil {
		return jwts, mistakes.NewError("构建jwts失败", err)
	}
	return jwts, nil
}

/*
* @desc 解析jwts
* 先判断jwt能否被正常解析
 */

func ParseJwt(secret string, jwts string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(jwts, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, mistakes.NewError("解析jwts失败", err)
	}
	return token, nil
}

func GetJwtsClamis(secret string, jwts string) (*CustomClaims, error) {
	token, err := ParseJwt(secret, jwts)
	if err != nil {
		return nil, mistakes.NewError("解析clamis失败,无法读取参数", err)
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, mistakes.NewError("Claims 结构存在问题", err)
	}
}
