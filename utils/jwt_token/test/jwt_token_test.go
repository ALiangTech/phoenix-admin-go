package test

import (
	jwtToken "phoenix-go-admin/utils/jwt_token"
	"testing"
)

var secret = "suuej^24**&"
var uuid = "234234"

// 测试jwts 生成

func TestGenerateJwts(t *testing.T) {

	jwts, _ := jwtToken.GenerateJwt(secret, uuid)
	claims, _ := jwtToken.GetJwtsClamis(secret, jwts)

	if claims.Uuid != uuid {
		t.Error("jwts 生成失败")
	} else {
		t.Log("jwts 生成成功")
	}

}

// 函数基准测试

func BenchmarkGenerateJwt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jwts, _ := jwtToken.GenerateJwt(secret, uuid)
		b.Log(jwts)
	}
}

func BenchmarkGetJwtsClamis(b *testing.B) {
	jwts, _ := jwtToken.GenerateJwt(secret, "234234")
	for i := 0; i < b.N; i++ {
		claims, _ := jwtToken.GetJwtsClamis(secret, jwts)
		b.Log(claims.Uuid)
	}
}
