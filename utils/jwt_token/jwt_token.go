package jwtToken

import (
	"phoenix-go-admin/utils/mistakes"

	"github.com/golang-jwt/jwt/v5"
)

/*
* @desc 生成jwts
* 支持自定义参数
 */

type customeClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateJwt(secret string, uuid string) (string, error) {
	// Create the Claims
	claims := customeClaims{
		uuid,
		jwt.RegisteredClaims{
			Issuer: "phoenix",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwts, err := token.SignedString([]byte(secret))
	return jwts, mistakes.NewError("构建jwts失败", err)
}

/*
* @desc 解析jwts
* 先判断jwt能否被正常解析
 */

func ParseJwt(secret string, jwts string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(jwts, &customeClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, mistakes.NewError("解析jwts失败", err)
	}
	return token, nil
}

/*
* @desc 从jwts 中获取自定义的参数
 */
func GetJwtsClamis(secret string, jwts string) (*customeClaims, error) {
	token, err := ParseJwt(secret, jwts)
	if err != nil {
		return nil, mistakes.NewError("解析clamis失败,无法读取参数", err)
	} else if claims, ok := token.Claims.(*customeClaims); ok {
		return claims, nil
	} else {
		return nil, mistakes.NewError("Claims 结构存在问题", err)
	}
}
