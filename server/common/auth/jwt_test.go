package auth_test

import (
	"beluga/server/common/auth"
	"reflect"
	"testing"
	"time"
)

func TestAuthToken(t *testing.T) {
	auth.InitJWTAuth("abcdef", time.Hour, time.Hour*24, "wangyc")
	atoken, err := auth.GetJWTAuth()
	if err != nil {
		t.Fatalf("token get failed")
	}
	claims := auth.UserClaims{
		UserId:       3861,
		UserName:     "wangyc",
		RealUserId:   3861,
		RealUserName: "wangyc",
		TokenType:    auth.AUTH,
	}
	token, err := atoken.NewToken(claims)
	if err != nil {
		t.Fatalf("new token failed: %v", err)
	}
	validteClaims, err := atoken.ValidateToken(token)
	if err != nil {
		t.Fatalf("validate token failed: %v", err)
	}
	if !reflect.DeepEqual(claims, validteClaims) {
		t.Fatalf("want: %v, got: %v", claims, validteClaims)
	}
}
