package users

import (
	"fmt"
	"github.com/sial-soft/users-api/datasources/postgres/users_db"
	"github.com/sial-soft/users-api/logger"
	"github.com/sial-soft/users-api/utils/errors"
	"github.com/sial-soft/users-api/utils/pg_utils"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryGetUser          = "SELECT id, first_name, last_name, email, created_at, status FROM users_db.users WHERE id=$1"
	queryInsertUser       = "INSERT INTO users_db.users(first_name, last_name, email, status, password) VALUES($1, $2, $3, $4, $5) RETURNING id"
	queryUpdateUser       = "UPDATE users_db.users SET first_name=$1, last_name=$2, email=$3, status=$4, password=$5 WHERE id=$6"
	queryDeleteUser       = "DELETE FROM users_db.users WHERE id=$1"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, created_at, status FROM users_db.users WHERE status=$1"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.NewInternalError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreateAt, &u.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return pg_utils.ParseError(err)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while trying to prepare save user statement", err)
		return errors.NewInternalError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.FirstName, u.LastName, u.Email, u.Status, u.Password)

	if err := row.Scan(&u.Id); err != nil {
		logger.Error("error when trying to save user", err)
		return pg_utils.ParseError(err)
	}

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while trying to prepare update user statement", err)
		return errors.NewInternalError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.Id)

	if err != nil {
		logger.Error("error when trying to update user", err)
		return pg_utils.ParseError(err)
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error while trying to prepare delete user statement", err)
		return errors.NewInternalError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)

	if err != nil {
		logger.Error("error when trying to delete user", err)
		return pg_utils.ParseError(err)
	}

	return nil
}

func (u *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error while trying to prepare search users statement", err)
		return nil, errors.NewInternalError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error while querying users", err)
		return nil, errors.NewInternalError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreateAt, &user.Status); err != nil {
			logger.Error("error when trying to scan user row", err)
			return nil, pg_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user found with status : %s", status))
	}
	return results, nil
}
