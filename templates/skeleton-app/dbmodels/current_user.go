package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
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

