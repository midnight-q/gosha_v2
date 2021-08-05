package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Code   string
	TypeId int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (resource *Resource) Validate() {
}
