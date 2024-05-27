package service

import (
	"ginfx-template/internal/model"
)

type IUserService interface {
	FindByUsername(username string) (model.User, error)
}

type UserService struct {
}

func (u *UserService) FindByUsername(username string) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService() IUserService {
	return &UserService{}
}

var _ IUserService = (*UserService)(nil)
