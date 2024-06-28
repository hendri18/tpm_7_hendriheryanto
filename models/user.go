package models

import (
	"tpm_7_HendriHeryanto/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"not null;uniqueIndex" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	hashedPassword, errHash := helpers.HasPass(u.Password)
	u.Password = hashedPassword

	if errHash != nil {
		err = errHash
		return
	}

	err = nil
	return
}
