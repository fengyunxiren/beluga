package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenType int

const (
	AUTH TokenType = iota + 1
	REFRESH
)

type UserClaims struct {
	UserId       int       `json:"user_id"`
	UserName     string    `json:"username"`
	RealUserId   int       `json:"real_user_id"`
	RealUserName string    `json:"real_user_name"`
	TokenType    TokenType `json:"token_type"`
}

type Claims struct {
	jwt.StandardClaims
	UserClaims
}

type JWTAuth struct {
	SecretKey   string
	ExpireTime  time.Duration
	RefreshTime time.Duration
	Issuer      string
}

func (j JWTAuth) CreateToken(userClaims UserClaims) (string, error) {
	var expiresAt time.Duration
	switch userClaims.TokenType {
	case AUTH:
		expiresAt = j.ExpireTime
	case REFRESH:
		expiresAt = j.RefreshTime
	default:
		return "", errors.New("Unsupported token type")
	}
	jwtClaims := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,
		ExpiresAt: time.Now().Add(expiresAt).Unix(),
	}
	claims := Claims{
		StandardClaims: jwtClaims,
		UserClaims:     userClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}

func (j JWTAuth) ValidateToken(tokenStr string) (*UserClaims, error) {
	var claims *Claims
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return j.SecretKey, nil
	})
	if err != nil {
		return &UserClaims{}, err
	}
	if token != nil {
		if token.Valid {
			return &claims.UserClaims, nil
		}
		return nil, errors.New("Token is not valid")

	} else {
		return nil, errors.New("Token is not valid")
	}
}
