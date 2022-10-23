package v1

import "github.com/gin-gonic/gin"

type TestApp struct{}

var TestApi TestApp

func (t TestApp) GetRouter(root gin.IRouter) gin.IRouter {
	router := root.Group("/api/v1")
	return router
}

func (t TestApp) RegisterRouter(router gin.IRouter) {
	router.GET("/apiCheck", APICheck)
}
