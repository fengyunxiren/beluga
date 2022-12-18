package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenType int

const (
	AUTH TokenType = iota + 1
	REFRESH
)

type UserClaims struct {
	UserId       uint64    `json:"user_id"`
	UserName     string    `json:"username"`
	RealUserId   uint64    `json:"real_user_id"`
	RealUserName string    `json:"real_user_name"`
	TokenType    TokenType `json:"token_type"`
}

type Claims struct {
	jwt.StandardClaims
	UserClaims
}

type JWTAuth struct {
	secretKey   string
	expireTime  time.Duration
	refreshTime time.Duration
	issuer      string
}

var defaultJWTAuth *JWTAuth

func GetJWTAuth() (*JWTAuth, error) {
	if defaultJWTAuth == nil {
		return nil, errors.New("JWT Auth is not found")
	} else {
		return defaultJWTAuth, nil
	}
}

func InitJWTAuth(secretKey string, expireTime, refreshTime time.Duration, issuer string) {
	jwt := JWTAuth{
		secretKey:   secretKey,
		expireTime:  expireTime,
		refreshTime: refreshTime,
		issuer:      issuer,
	}
	defaultJWTAuth = &jwt
}

func (j JWTAuth) NewToken(userClaims UserClaims) (string, error) {
	var expiresAt time.Duration
	switch userClaims.TokenType {
	case AUTH:
		expiresAt = j.expireTime
	case REFRESH:
		expiresAt = j.refreshTime
	default:
		return "", errors.New("Unsupported token type")
	}
	jwtClaims := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,
		ExpiresAt: time.Now().Add(expiresAt).Unix(),
		Issuer:    j.issuer,
		Id:        uuid.NewString(),
	}
	claims := Claims{
		StandardClaims: jwtClaims,
		UserClaims:     userClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j JWTAuth) ValidateToken(tokenStr string) (UserClaims, error) {
	var claims Claims = Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return UserClaims{}, err
	}
	if token != nil {
		if token.Valid {
			return claims.UserClaims, nil
		}
		return UserClaims{}, errors.New("Token is not valid")

	} else {
		return UserClaims{}, errors.New("Token is not valid")
	}
}
