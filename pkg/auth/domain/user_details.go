package domain

import "time"

type UserDetails interface {
	GetAuthorities() []string
	GetPassword() string
	GetName() string
	GetUsername() string
	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool
	IsEnabled() bool
	GetLastLogin() time.Time
}

type User struct {
	Name                  string
	Username              string
	Password              string
	Authorities           []string
	AccountNonExpired     bool
	AccountNonLocked      bool
	CredentialsNonExpired bool
	Enabled               bool
	LastLogin             time.Time
}

var _ UserDetails = (*User)(nil)

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAuthorities() []string {
	return u.Authorities
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) IsAccountNonExpired() bool {
	return u.AccountNonExpired
}

func (u *User) IsAccountNonLocked() bool {
	return u.AccountNonLocked
}

func (u *User) IsCredentialsNonExpired() bool {
	return u.CredentialsNonExpired
}

func (u *User) IsEnabled() bool {
	return u.Enabled
}

func (u *User) GetLastLogin() time.Time {
	return u.LastLogin
}
