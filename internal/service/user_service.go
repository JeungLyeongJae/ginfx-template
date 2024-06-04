package service

import (
	"ginfx-template/internal/model"
	"ginfx-template/internal/repository"
)

type IUserService interface {
	Save(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

type UserService struct {
	userRepo repository.IUserRepo
}

func NewUserService(userRepo repository.IUserRepo) IUserService {
	return &UserService{userRepo: userRepo}
}

var _ IUserService = (*UserService)(nil)

func (u *UserService) Save(user *model.User) error {
	return u.userRepo.Save(user)
}

func (u *UserService) FindByUsername(username string) (*model.User, error) {
	return u.userRepo.FindByUsername(username)
}
