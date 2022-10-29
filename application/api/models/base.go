package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time      `json:"create_at,omitempty"`
	UpdatedAt time.Time      `json:"update_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"delete_at,omitempty" gorm:"index"`
}
