package model

import (
	"time"
)

type User struct {
	Id            int64      `json:"id" gorm:"primary_key; column:id; comment:用户ID"`
	Name          string     `json:"name" gorm:"column:name; type:varchar(255); comment:名字"`
	Phone         string     `json:"phone" gorm:"column:phone; type:varchar(255); comment:手机号"`
	Email         string     `json:"email" gorm:"column:email; type:varchar(255); comment:邮箱"`
	Username      string     `json:"username" gorm:"column:username; type:varchar(255); comment:登陆账号名称"`
	Password      string     `json:"password" gorm:"column:password; type:varchar(255); comment:登陆账号密码（明文）"`
	PlainPassword string     `json:"plain_password" gorm:"column:plain_password; type:varchar(255); comment:登陆账号密码（密文）"`
	Roles         []string   `json:"roles" gorm:"column:roles; type:varchar(255); comment:用户权限"`
	Enable        bool       `json:"enable" gorm:"column:enable; type:tinyint(1); comment:是否禁用（0： 被禁用； 1： 启用）"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at; comment:创建时间; autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"column:updated_at; comment:更新时间; autoUpdateTime"`
	LastLogin     *time.Time `json:"last_login" gorm:"column:last_login; comment:最后一次登陆时间"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"column:deleted_at; comment:删除时间;"`
}

func (User) TableName() string {
	return "auth_user"
}
