package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (role *Role) Validate() {
}
