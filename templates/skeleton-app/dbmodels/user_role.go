package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type UserRole struct {

    ID        int       `gorm:"primary_key"`
    UserId int
	RoleId int

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (userRole *UserRole) Validate() {
}

