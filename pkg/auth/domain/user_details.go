package domain

type UserDetails interface {
	GetAuthorities() []string
	GetPassword() string
	GetName() string
	GetUsername() string
	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool
	IsEnabled() bool
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
