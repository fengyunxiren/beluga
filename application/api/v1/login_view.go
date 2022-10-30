package v1

import (
	"beluga/application/api/models"
	"beluga/application/api/v1/forms"
	"beluga/server/common/database"
	"beluga/server/common/logger"
	"beluga/server/common/response"
	"beluga/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UserLogin(c *gin.Context) {
	log := logger.GetContextLogger(c)
	db, err := database.GetContextDB(c)
	if err != nil {
		log.Error("Get db Failed: ", err)
		response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
		return
	}
	form := forms.LoginForm{}
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Error("Bind body failed: ", err)
		response.AbortWithError(response.ERR_BAD_REQUEST, c)
		return
	}
	user := models.User{}
	result := db.First(&user, "user_name=?", form.UserName)
	if result.Error != nil {
		log.Error("Login failed: ", result.Error)
		response.ResponseError(response.ERR_USER_OR_PASSWORD, c)
		return
	}
	ok := utils.ValidatePassword(user.Password, form.Password)
	if !ok {
		response.ResponseError(response.ERR_USER_OR_PASSWORD, c)
		return
	}
	// TODO, return Token
	response.ResponseOkWithData(gin.H{"token": "token"}, c)
}
