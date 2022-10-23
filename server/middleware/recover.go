package middleware

import (
	"beluga/server/common/logger"
	"beluga/server/common/response"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler(c *gin.Context, err interface{}) {
	log := logger.GetContextLogger(c)
	log.Error("service return error: ", err)
	response.AbortWithError(response.ERR_SERVER_500, c)
}

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(RecoveryHandler)
}
