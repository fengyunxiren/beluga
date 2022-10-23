package middleware

import (
	"beluga/global"
	"beluga/server/common/database"
	"beluga/server/common/response"

	"github.com/gin-gonic/gin"
)

func SetDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		if db == nil {
			response.AbortWithError(response.ERR_SERVER_500, c)
		}
		c.Set(global.CONTEXT_DB, db)
	}
}
