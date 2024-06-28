package service

import (
	"tpm_7_HendriHeryanto/models"
	"tpm_7_HendriHeryanto/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func (p *UserService) GetByEmail(email string) (*models.User, error) {
	return p.UserRepo.GetUserByEmail(email)
}

func (p *UserService) Create(user *models.User) (*models.User, error) {
	return p.UserRepo.CreateUser(user)
}
