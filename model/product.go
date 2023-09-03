package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
