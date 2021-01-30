package users

import (
	"fmt"
	"github.com/sial-soft/users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.CreateAt = result.CreateAt
	return nil
}

func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]
	if current != nil {
		return errors.NewBadRequest(fmt.Sprintf("user %d alreay exists", u.Id))
	}
	usersDB[u.Id] = u
	return nil
}
