package v1

import (
	"beluga/server/common/response"

	"github.com/gin-gonic/gin"
)

func APICheck(c *gin.Context) {
	response.ResponseOkWithData(
		gin.H{
			"name": "gin",
		}, c,
	)
}
