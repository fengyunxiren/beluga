package server

import (
	"beluga/application/api/models"
	"beluga/utils"
)

var DBGenerator []utils.Generator = []utils.Generator{
	utils.NewGenerator[models.User](),
}
