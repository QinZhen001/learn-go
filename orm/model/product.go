package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

type Food struct {
	FoodId uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:food_name;type:varchar(32);index:idx_food_name,unique"`
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
