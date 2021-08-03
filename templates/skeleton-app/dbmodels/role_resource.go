package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type RoleResource struct {

    ID        int       `gorm:"primary_key"`
    RoleId int
	ResourceId int
	Find bool
	Read bool
	Create bool
	Update bool
	Delete bool
	FindOrCreate bool
	UpdateOrCreate bool
	//RoleResource remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (roleResource *RoleResource) Validate() {
    //Validate remove this line for disable generator functionality
}

