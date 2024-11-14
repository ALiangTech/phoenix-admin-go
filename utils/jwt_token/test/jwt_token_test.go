package test

import (
	jwtToken "phoenix-go-admin/utils/jwt_token"
	"testing"
)

// 测试jwts 生成

func TestGenerateJwts(t *testing.T) {
	var secret = "suuej^24**&"
	jwts, _ := jwtToken.GenerateJwt(secret, "234234")
	t.Log(jwts)
	claims, _ := jwtToken.GetJwtsClamis(secret, jwts)
	t.Log(claims.Uuid)

}
