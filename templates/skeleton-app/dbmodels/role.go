package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type Role struct {

    ID        int       `gorm:"primary_key"`
    Name string
	Description string
	//Role remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (role *Role) Validate() {
    //Validate remove this line for disable generator functionality
}

