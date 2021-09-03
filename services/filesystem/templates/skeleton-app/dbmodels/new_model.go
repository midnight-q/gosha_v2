package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type NewModel struct {
	ID int `gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

	validator
}

func (model *NewModel) Validate() {
}
