package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64         `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time      `json:"create_at,omitempty"`
	UpdatedAt time.Time      `json:"update_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
