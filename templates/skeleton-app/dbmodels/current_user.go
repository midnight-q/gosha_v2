package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type CurrentUser struct {

    ID        int       `gorm:"primary_key"`
    //CurrentUser remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (currentUser *CurrentUser) Validate() {
    //Validate remove this line for disable generator functionality
}

