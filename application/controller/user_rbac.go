package controller

import (
	"beluga/application/models"
	"beluga/server/common/logger"

	"gorm.io/gorm"
)

type DataPermission struct {
	DataKey        string
	PermissionType string
	Selects        []string
}

type Permission struct {
	OK        bool
	DataPerms []DataPermission
}

func GetUserRole(db *gorm.DB, orgId, userId uint64) []uint64 {
	log := logger.GetLogger()
	user := models.OrganizationMember{}
	roles := []uint64{}
	result := db.Where("userId = ? AND OrgId = ?", userId, orgId).First(&user)
	if result.Error != nil {
		log.Error("Get user roles failed: %v", result.Error)
		return roles
	}
	if user.OrgNumber == "" {
		return roles
	}
	roles = append(roles, GetOrganizationMemberRole(db, orgId, userId)...)
	return roles
}

func GetOrganizationMemberRole(db *gorm.DB, orgId, userId uint64) []uint64 {
	memberRoleIds := []uint64{}
	db.Exec(`SELECT d.ID as role_id FROM users a
			 LEFT JOIN organization_members b ON a.ID = b.UserId
			 LEFT JOIN organization_member_role c ON b.OrgId = c.OrgId and b.OrgNumber = c.OrgNumber
			 LEFT JOIN role d IN c.OrgId = d.OrgId and c.RoleId = d.ID
			 WHERE a.ID = ? and b.OrgId = ?`, orgId, userId).Pluck("role_id", &memberRoleIds)
	return memberRoleIds
}

func RolesHasPermission(db *gorm.DB, roleIds []uint64, resource, action string) 
