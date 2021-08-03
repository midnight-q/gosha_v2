package dbmodels

import (
    "time"

    "gorm.io/gorm"
)

type CurrentUser struct {

    ID        int       `gorm:"primary_key"`

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (currentUser *CurrentUser) Validate() {
}

