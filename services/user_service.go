package services

import (
	"github.com/sial-soft/users-api/domain/users"
	"github.com/sial-soft/users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if user.Email == "" {
		return nil, errors.NewInternalError("database is down")
	}
	return &user, nil
}
