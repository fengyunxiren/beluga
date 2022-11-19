package auth

import (
	"beluga/application/api/models"
	"beluga/application/api/v1/forms"
	"beluga/server/common/database"
	"beluga/server/common/logger"
	"beluga/server/common/response"
	"beluga/utils"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UserRegister(c *gin.Context) {
	log := logger.GetContextLogger(c)
	db, err := database.GetContextDB(c)
	if err != nil {
		log.Error("Get db Failed: ", err)
		response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
		return
	}
	form := forms.RegisterForm{}
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Error("Bind body failed: ", err)
		response.AbortWithError(response.ERR_BAD_REQUEST, c)
		return
	}
	hashPassword, err := utils.HashPassword(form.Password)
	if err != nil {
		log.Error("Password encrypt failed: ", err)
		response.AbortWithError(response.ERR_SERVER_500, c)
		return
	}
	form.Password = hashPassword
	user := models.User{}
	result := db.Model(&user).Select("UserName", "Password", "Email").Create(
		// TODO, Struct To Map
		map[string]interface{}{
			"UserName": form.UserName,
			"Password": form.Password,
			"Email":    form.Email,
		},
	)
	if result.Error != nil {
		log.Error("Create User failed: ", result.Error)
		fmt.Println("error: ", reflect.TypeOf(result.Error))
		if utils.IsDuplicateKeyError(result.Error) {
			response.ResponseError(response.ERR_REIGSTER_UNIQUE, c)
		} else {
			response.ResponseError(response.ERR_SERVER_VIEW_CREATE, c)
		}
	} else {
		response.ResponseOk(c)
	}
}
