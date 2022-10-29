package v1

import (
	"beluga/application/api/models"
	"beluga/server/common/view"
	"beluga/utils"

	"github.com/gin-gonic/gin"
)

type TestAPIRouter struct{}

var TestAPI TestAPIRouter

func (t TestAPIRouter) GetRouter(root gin.IRouter) gin.IRouter {
	router := root.Group("/api/v1")
	return router
}

func (t TestAPIRouter) RegisterRouter(router gin.IRouter) {
	router.GET("/apiCheck", APICheck)
	userView := view.NewDBView("", "User", utils.NewGenerator[models.User]())
	view.RegisterView(userView, router)
}
