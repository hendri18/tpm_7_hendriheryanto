package repository

import (
	"errors"
	"tpm_7_HendriHeryanto/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := u.DB.Debug().Where("email = ?", email).Find(&user)

	err := result.Error

	if result.RowsAffected < 1 {
		err = errors.New("user tidak ditemukan")
	}

	return user, err
}

func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	err := u.DB.Debug().Create(&user).Error
	return user, err
}
