package response

import (
	"beluga/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrorCode
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

func GetRequestId(c *gin.Context) string {
	var requestId string
	context, exists := c.Get(global.CONTEXT_REQUEST_ID)
	if !exists {
		requestId = ""
	} else {
		strContext, ok := context.(string)
		if !ok {
			requestId = ""
		} else {
			requestId = strContext
		}
	}
	return requestId
}

func NewResponse(code ErrorCode, data interface{}, requestId string) Response {
	response := Response{
		ErrorCode: code,
		Data:      data,
		RequestId: requestId,
	}
	return response
}

func ResponseWithStatus(status int, code ErrorCode, data interface{}, c *gin.Context) {
	requestId := GetRequestId(c)
	response := NewResponse(code, data, requestId)
	c.JSON(status, response)
}

func ResponseWithData(code ErrorCode, data interface{}, c *gin.Context) {
	ResponseWithStatus(http.StatusOK, code, data, c)
}

func ResponseOkWithData(data interface{}, c *gin.Context) {
	ResponseWithData(OK, data, c)
}

func ResponseOk(c *gin.Context) {
	ResponseOkWithData(nil, c)
}

func ResponseError(code ErrorCode, c *gin.Context) {
	ResponseWithData(code, nil, c)
}

func AbortWithError(code ErrorCode, c *gin.Context) {
	requestId := GetRequestId(c)
	response := NewResponse(code, nil, requestId)
	c.AbortWithStatusJSON(http.StatusOK, response)
}
