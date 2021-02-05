package users

import (
	"github.com/sial-soft/users-api/utils/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  string `json:"create_at"`
}

func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}
