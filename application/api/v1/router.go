package v1

import (
	"github.com/gin-gonic/gin"
)

type APIV1Router struct{}

var APIV1 APIV1Router

func (t APIV1Router) GetRouter(root gin.IRouter) gin.IRouter {
	router := root.Group("/api/v1")
	return router
}

func (t APIV1Router) RegisterRouter(router gin.IRouter) {
	router.POST("/users/register", UserRegister)
}
