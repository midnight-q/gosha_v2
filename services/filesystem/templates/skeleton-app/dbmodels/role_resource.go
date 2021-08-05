package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type RoleResource struct {
	ID             int `gorm:"primary_key"`
	RoleId         int
	ResourceId     int
	Find           bool
	Read           bool
	Create         bool
	Update         bool
	Delete         bool
	FindOrCreate   bool
	UpdateOrCreate bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (roleResource *RoleResource) Validate() {
}
