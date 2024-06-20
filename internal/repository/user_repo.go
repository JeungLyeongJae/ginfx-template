package repository

import (
	"ginfx-template/internal/model"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindByUsername(name string) (*model.User, error)
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

func (u *UserRepo) FindByUsername(name string) (*model.User, error) {
	user := &model.User{}
	if err := u.db.First(user, "username = ?", name).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Save(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) Delete(user *model.User) error {
	return u.db.Delete(user).Error
}

func (u *UserRepo) Update(user *model.User) error {
	return u.db.Save(user).Error
}

var _ IUserRepo = (*UserRepo)(nil)
