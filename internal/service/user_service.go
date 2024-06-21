package service

import (
	"ginfx-template/internal/model"
	"ginfx-template/internal/model/response"
	"ginfx-template/internal/repository"
)

type IUserService interface {
	GetUserList(*response.Page) error
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

func (u *UserService) GetUserList(page *response.Page) error {
	return u.userRepo.GetUserList(page)
}

func (u *UserService) Save(user *model.User) error {
	return u.userRepo.Save(user)
}

func (u *UserService) FindByUsername(username string) (*model.User, error) {
	return u.userRepo.FindByUsername(username)
}
