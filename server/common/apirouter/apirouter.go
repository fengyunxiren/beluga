package apirouter

import "github.com/gin-gonic/gin"

type IAPIRouter interface {
	GetRouter(root gin.IRouter) gin.IRouter
	RegisterRouter(router gin.IRouter)
}

type IAPIRouters []IAPIRouter

func (apis IAPIRouters) RegisterRouter(root gin.IRouter) {
	for _, api := range apis {
		router := api.GetRouter(root)
		api.RegisterRouter(router)
	}
}
