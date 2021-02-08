package pg_utils

import (
	"github.com/lib/pq"
	"github.com/sial-soft/users-api/utils/errors"
	"strings"
)

const (
	noRowError = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	pgErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), noRowError) {
			return errors.NewNotFoundError("no record found matching id")
		}
		return errors.NewInternalError("error parsing database response")
	}
	switch pgErr.Code {
	case "23505":
		return errors.NewBadRequest("unique constraint violation: duplicated data")
	}
	return errors.NewInternalError("error posting request")
}
