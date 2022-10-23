package middleware

import (
	"/logger"
	"kangaroo/common/response"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler(c *gin.Context, err interface{}) {
	log := logger.GetContextEntry(c)
	log.Error("service return error: ", err)
	response.AbortWithError(response.ERR_SERVER_500, c)
}

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(RecoveryHandler)
}
