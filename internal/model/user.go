package model

import "time"

type User struct {
	Id            int64
	Name          string
	Phone         string
	Username      string
	Password      string
	PlainPassword string
	Roles         []string
	Created       time.Time
	UpdatedAt     time.Time
	Enable        bool
	LastLogin     time.Time
	DeletedAt     time.Time
}
