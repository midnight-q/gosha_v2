package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type Language struct {

    ID        int       `gorm:"primary_key"`
    Name int
	Code string
	//Language remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (language *Language) Validate() {
    //Validate remove this line for disable generator functionality
}

