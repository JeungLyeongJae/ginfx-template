package services

import (
	"ginfx-template/internal/service"
	"ginfx-template/pkg/auth/domain"
)

type UserDetailsService interface {
	LoadUserByUsername(username string) (domain.UserDetails, error)
	UpdateLastLogin(username string) error
}

type DbUserDetailsService struct {
	service.IUserService
}

func (m *DbUserDetailsService) UpdateLastLogin(username string) error {
	if err := m.UpdateLastLoginTime(username); err != nil {
		return err
	}
	return nil
}

var _ UserDetailsService = (*DbUserDetailsService)(nil)

func NewUserDetailsService(userService service.IUserService) UserDetailsService {
	return &DbUserDetailsService{
		userService,
	}
}

func (m *DbUserDetailsService) LoadUserByUsername(username string) (domain.UserDetails, error) {
	u, err := m.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Name:                  u.Name,
		Username:              u.Username,
		Password:              u.PlainPassword,
		Authorities:           u.Roles,
		AccountNonExpired:     true,
		AccountNonLocked:      true,
		CredentialsNonExpired: true,
		Enabled:               true,
	}, nil
}
