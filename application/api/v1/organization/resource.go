package organization

import (
	"beluga/application/api/models"
	"fmt"

	"gorm.io/gorm"
)

type OrganizationResource struct {
	db *gorm.DB
}

func (t OrganizationResource) ListOrganization(name, code string) ([]models.Organization, error) {
	result := []models.Organization{}
	m := models.Organization{}

	exec := t.db.Model(&m)
	if name != "" {
		exec = exec.Where("Name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	if code != "" {
		exec = exec.Where("Code = ?", code)
	}
	exec.Find(&result)
	
}
