package services

import (
	"ginfx-template/internal/service"
	"ginfx-template/pkg/auth/domain"
)

type UserDetailsService interface {
	LoadUserByUsername(username string) (domain.UserDetails, error)
}

type DbUserDetailsService struct {
	service.IUserService
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
