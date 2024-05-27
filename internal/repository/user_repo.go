package repository

import (
	"ginfx-template/internal/model"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindByName(name string) (*model.User, error)
	Save(*model.User) error
	Delete(*model.User) error
	Update(*model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindByName(name string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Save(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

var _ IUserRepo = (*UserRepo)(nil)
