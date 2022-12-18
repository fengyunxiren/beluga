package admin

import (
	"beluga/application/models"
	"beluga/server/common/view"
	"beluga/utils"
)

var AdminViews []view.View = []view.View{
	view.NewDBView("", "users", utils.NewGenerator[models.User]()),
	view.NewDBView("", "organizations", utils.NewGenerator[models.Organization]()),
	view.NewDBView("", "organization/members", utils.NewGenerator[models.OrganizationMember]()),
	view.NewDBView("", "organization/departments", utils.NewGenerator[models.Department]()),
	view.NewDBView("", "organization/department/members", utils.NewGenerator[models.DepartmentMember]()),
	view.NewDBView("", "organization/groups", utils.NewGenerator[models.OrganizationGroup]()),
	view.NewDBView("", "organization/group/members", utils.NewGenerator[models.OrganizationGroupMember]()),
}
