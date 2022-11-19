package middleware

import (
	"beluga/global"
	"beluga/server/common/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(global.AUTH_HEADER)
		fmt.Println("token: ", token)
		if token == "" {
			c.AbortWithStatus(401)
			return
		}
		jwtAuth, err := auth.GetJWTAuth()
		if err != nil {
			fmt.Println("error: ", err)
			c.AbortWithStatus(401)
			return
		}
		claims, err := jwtAuth.ValidateToken(token)
		if err != nil {
			fmt.Println("validate error: ", err)
			c.AbortWithStatus(401)
			return
		}
		fmt.Println("clams: ", claims)
		c.Next()
	}
}
