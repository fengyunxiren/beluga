package organization

import (
	"beluga/server/common/database"
	"beluga/server/common/logger"

	"github.com/gin-gonic/gin"
)

func ListOrganization(c *gin.Context) {
	log := logger.GetContextLogger(c)
	db, err := database.GetContextDB(c)
}
