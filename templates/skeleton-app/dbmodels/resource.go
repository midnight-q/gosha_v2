package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type Resource struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Code string
	TypeId int

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (resource *Resource) Validate() {
}

