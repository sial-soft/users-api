package users

import (
	"github.com/sial-soft/users-api/datasources/postgres/users_db"
	"github.com/sial-soft/users-api/utils/errors"
	"github.com/sial-soft/users-api/utils/pg_utils"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users_db.users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING id"
	queryUpdateUser = "UPDATE users_db.users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4"
	queryDeleteUser = "DELETE FROM users_db.users WHERE id=$1"
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
		return pg_utils.ParseError(err)
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
		return pg_utils.ParseError(err)
	}

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)

	if err != nil {
		return pg_utils.ParseError(err)
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)

	if err != nil {
		return pg_utils.ParseError(err)
	}

	return nil
}
