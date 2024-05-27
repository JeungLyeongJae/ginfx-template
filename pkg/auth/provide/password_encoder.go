package provide

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type PasswordEncoder interface {
	Encode(rawPassword string) (string, error)
	Matches(rawPassword string, encodedPassword string) bool
	Prefix() string
}

var Encoder = &BCryptPasswordEncoder{}

type BCryptPasswordEncoder struct {
}

var _ PasswordEncoder = (*BCryptPasswordEncoder)(nil)

func (b *BCryptPasswordEncoder) Encode(rawPassword string) (string, error) {
	encodePw, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	} else {
		return b.Prefix() + string(encodePw), nil
	}
}

func (b *BCryptPasswordEncoder) Matches(rawPassword, passWd string) bool {
	raw := strings.TrimPrefix(passWd, b.Prefix())
	return bcrypt.CompareHashAndPassword([]byte(raw), []byte(rawPassword)) == nil
}

func (b *BCryptPasswordEncoder) Prefix() string {
	return "{bcrypt}"
}
