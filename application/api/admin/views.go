package admin

import (
	"beluga/application/api/models"
	"beluga/server/common/view"
	"beluga/utils"
)

var AdminViews []view.View = []view.View{
	view.NewDBView("", "user", utils.NewGenerator[models.User]()),
}
