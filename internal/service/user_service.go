package service

import (
	"errors"
	"fmt"
	"ginfx-template/internal/model"
	"ginfx-template/internal/model/response"
	"ginfx-template/internal/repository"
	"ginfx-template/pkg/common"
)

type IUserService interface {
	IsExistByUsername(username string) (bool, error)
	GetUserList(*response.Page) error
	Create(user *model.User) error
	Update(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	Delete(user *model.User) error
	UpdateLastLoginTime(username string) error
}

type UserService struct {
	userRepo repository.IUserRepo
}

func NewUserService(userRepo repository.IUserRepo) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

var _ IUserService = (*UserService)(nil)

func (u *UserService) IsExistByUsername(username string) (bool, error) {
	return u.userRepo.IsExistByUsername(username)
}

func (u *UserService) GetUserList(page *response.Page) error {
	return u.userRepo.GetUserList(page)
}

func (u *UserService) Create(user *model.User) error {
	isExist, err := u.IsExistByUsername(user.Username)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(fmt.Sprintf(`[%s] 用户名已存在！请重新创建用户！`, user.Username))
	}
	user.PlainPassword, err = common.Encoder.Encode(user.Password)
	if err != nil {
		return err
	}
	return u.userRepo.Create(user)
}

func (u *UserService) Update(user *model.User) error {
	encode, err := common.Encoder.Encode(user.Password)
	if err != nil {
		return err
	}
	user.PlainPassword = encode
	return u.userRepo.Update(user)
}

func (u *UserService) FindByUsername(username string) (*model.User, error) {
	return u.userRepo.FindByUsername(username)
}

func (u *UserService) Delete(user *model.User) error {
	return u.userRepo.Delete(user)
}

func (u *UserService) UpdateLastLoginTime(username string) error {
	return u.userRepo.UpdateLastLogin(username)
}
