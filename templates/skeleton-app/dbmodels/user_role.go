package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type UserRole struct {

    ID        int       `gorm:"primary_key"`
    UserId int
	RoleId int
	//UserRole remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (userRole *UserRole) Validate() {
    //Validate remove this line for disable generator functionality
}

