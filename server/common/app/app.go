package app

import "github.com/gin-gonic/gin"

type IApp interface {
	GetRouter(root gin.IRouter) gin.IRouter
	RegisterRouter(router gin.IRouter)
}

type IApps []IApp

func (apps IApps) RegisterRouter(root gin.IRouter) {
	for _, app := range apps {
		router := app.GetRouter(root)
		app.RegisterRouter(router)
	}
}
