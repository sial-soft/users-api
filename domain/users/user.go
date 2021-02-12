package users

import (
	"github.com/sial-soft/users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  string `json:"create_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

func (u *User) Validate(partial bool) *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)
	u.Password = strings.TrimSpace(u.Password)
	u.Status = strings.TrimSpace(u.Status)
	if partial && u.Email == "" {
		return nil
	} else if u.Email == "" || !strings.Contains(u.Email, "@") {
		return errors.NewBadRequest("invalid email address")
	}
	if partial && u.Password == "" {
		return nil
	} else if u.Password == "" || len(u.Password) < 5 {
		return errors.NewBadRequest("invalid password")
	}
	return nil
}
