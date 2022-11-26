package controller

import (
	"beluga/application/models"

	"gorm.io/gorm"
)

type PermissionController struct {
	db       *gorm.DB
	Resource string
	Action   string
	OrgId    uint64
}

func (p PermissionController) HasPermission(userId uint64) bool {
	user := models.OrganizationMember{}	
	p.db.Where("userId = ?", userId).First(&user)
}
