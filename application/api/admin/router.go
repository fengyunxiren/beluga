package admin

import (
	"beluga/server/common/view"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

var Admin AdminRouter

func (t AdminRouter) GetRouter(root gin.IRouter) gin.IRouter {
	router := root.Group("/api/admin")
	return router
}

func (t AdminRouter) RegisterRouter(router gin.IRouter) {
	for _, v := range AdminViews {
		view.RegisterView(v, router)
	}
}