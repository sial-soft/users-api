package users

import (
	"fmt"
	"github.com/sial-soft/users-api/datasources/postgres/users_db"
	"github.com/sial-soft/users-api/utils/errors"
	"strings"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users_db.users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING id"
	queryGetUser    = "SELECT id, first_name, last_name, email, created_at FROM users_db.users WHERE id=$1"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreateAt); err != nil {
		if strings.Contains(err.Error(), " no rows in result set") {
			return errors.NewNotFoundError(fmt.Sprintf("user not found with id: %d, %s", u.Id, err.Error()))
		}
		return errors.NewInternalError(fmt.Sprintf("error when getting user id: %d, %s", u.Id, err.Error()))
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.FirstName, u.LastName, u.Email)

	if err := row.Scan(&u.Id); err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return errors.NewBadRequest("same email already exist")
		}
		return errors.NewInternalError(fmt.Sprintf("error while trying to get id: %s", err.Error()))
	}

	return nil
}
