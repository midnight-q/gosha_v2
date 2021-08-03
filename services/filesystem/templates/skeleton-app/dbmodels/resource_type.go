package dbmodels

import (
    "time"

    "gorm.io/gorm"
)

type ResourceType struct {

    ID        int       `gorm:"primary_key"`
    Name string

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (resourceType *ResourceType) Validate() {
}

