package services

import (
	"github.com/sial-soft/users-api/domain/users"
	"github.com/sial-soft/users-api/utils/crypto_utils"
	"github.com/sial-soft/users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(userId int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(user users.User, partial bool) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	Search(status string) (users.Users, *errors.RestErr)
}

func (u *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(false); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) UpdateUser(user users.User, partial bool) (*users.User, *errors.RestErr) {
	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(partial); err != nil {
		return nil, err
	}

	if partial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Password != "" {
			current.Password = user.Password
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (u *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (u *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)

}
