package server

import (
	"beluga/application/models"
	"beluga/utils"
)

var DBGenerator []utils.Generator = []utils.Generator{
	utils.NewGenerator[models.User](),
	utils.NewGenerator[models.Organization](),
	utils.NewGenerator[models.OrganizationMember](),
	utils.NewGenerator[models.Department](),
	utils.NewGenerator[models.DepartmentMember](),
	utils.NewGenerator[models.OrganizationGroup](),
	utils.NewGenerator[models.OrganizationGroupMember](),
}
