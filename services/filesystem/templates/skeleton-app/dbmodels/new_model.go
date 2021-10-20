package dbmodels

import (
	"time"
)

type NewModel struct {
	ID int `gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time
	validator
}

func (model *NewModel) Validate() {
}
