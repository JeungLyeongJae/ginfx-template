package repository

import (
	"ginfx-template/internal/model"
	"ginfx-template/internal/model/response"
	"gorm.io/gorm"
	"time"
)

type IUserRepo interface {
	IsExistByUsername(username string) (bool, error)
	GetUserList(*response.Page) error
	FindByUsername(name string) (*model.User, error)
	Create(*model.User) error
	Delete(*model.User) error
	Update(*model.User) error
	UpdateLastLogin(username string) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) IsExistByUsername(username string) (bool, error) {
	var count int64
	u.db.Where("username = ?", username).First(&model.User{}).Count(&count)
	return count > 0, nil
}

func (u *UserRepo) GetUserList(page *response.Page) error {
	if page.Condition != "" {
		return u.db.Model(&model.User{}).
			Count(&page.TotalCount).
			Offset((page.PageNumber-1)*page.PageSize).
			Limit(page.PageSize).
			Find(&page.Users, "username like ?", "%"+page.Condition+"%").
			Error
	}
	return u.db.Model(&model.User{}).
		Count(&page.TotalCount).
		Offset((page.PageNumber - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&page.Users).
		Error
}

func (u *UserRepo) FindByUsername(name string) (*model.User, error) {
	user := &model.User{}
	if err := u.db.First(user, "username = ?", name).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Create(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) Delete(user *model.User) error {
	return u.db.Delete(user).Error
}

func (u *UserRepo) Update(user *model.User) error {
	return u.db.Model(user).Select("*").Omit("created_at").Updates(user).Error
}

func (u *UserRepo) UpdateLastLogin(username string) error {
	// skip Hooks methods and don’t track the update time when updating,
	// use UpdateColumn, UpdateColumns, it works like Update, Updates
	return u.db.Model(&model.User{}).Where("username = ?", username).UpdateColumn("last_login", time.Now()).Error
}

var _ IUserRepo = (*UserRepo)(nil)
