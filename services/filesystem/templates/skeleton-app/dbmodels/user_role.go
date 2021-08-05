package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID     int `gorm:"primary_key"`
	UserId int
	RoleId int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (userRole *UserRole) Validate() {
}
