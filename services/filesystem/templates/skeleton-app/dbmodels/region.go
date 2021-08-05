package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type Region struct {
	ID   int `gorm:"primary_key"`
	Name int
	Code string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (region *Region) Validate() {
}
