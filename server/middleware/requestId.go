package middleware

import (
	"beluga/global"
	"beluga/server/common/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		header := global.X_REQUEST_HEADER
		requestId := c.GetHeader(header)
		if requestId == "" {
			requestId = c.GetHeader(strings.ToLower(header))
		}
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Request.Header.Set(header, requestId)
		c.Set(global.CONTEXT_REQUEST_ID, requestId)
		fields := logrus.Fields{
			global.CONTEXT_REQUEST_ID: requestId,
		}
		log := logger.GetLogger()
		entry := log.WithFields(fields)
		c.Set(global.CONTEXT_LOGGER, entry)
		c.Next()
	}
}
