package dbmodels

import (
    "skeleton-app/common"
    "skeleton-app/errors"
    "time"

    "gorm.io/gorm"
)

type User struct {

    ID        int       `gorm:"primary_key"`
    Email       string  `gorm:"type:varchar(100);unique_index"`
    FirstName   string
    IsActive    bool
    LastName    string
    MobilePhone string
    Password    string

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `sql:"index" json:"-"`

    validator
}

func (user *User) Validate() {
    

    if len(user.FirstName) < 1 {
        user.AddValidationError("User first name is empty", errors.ErrorCodeFieldLengthTooShort, "FirstName")
    }

    if len(user.LastName) < 1 {
        user.AddValidationError("User last name is empty", errors.ErrorCodeFieldLengthTooShort, "LastName")
    }

    if len(user.Email) < 3 || ! common.ValidateEmail(user.Email)  {
        user.AddValidationError("User email not valid", errors.ErrorCodeNotValid, "Email")
    }

    if len(user.MobilePhone) > 3 &&  ! common.ValidateMobile(user.MobilePhone)  {
        user.AddValidationError("User mobile phone should be valid or empty. Format +0123456789... ", errors.ErrorCodeNotValid, "MobilePhone")
    }

}

