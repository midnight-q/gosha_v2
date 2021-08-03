package dbmodels

import (
    "gorm.io/gorm"
    "time"
    
)

type TranslateError struct {

    ID        int       `gorm:"primary_key"`
    Code int
	LanguageCode string
	Translate string
	//TranslateError remove this line for disable generator functionality

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (translateError *TranslateError) Validate() {
    //Validate remove this line for disable generator functionality
}

