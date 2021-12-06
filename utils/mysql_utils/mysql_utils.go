package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"go-apk-users/app/config"
	"go-apk-users/utils/errors"
)

func ParseErrors(err error) *errors.RestErr {
	sqlErr, _ := err.(*mysql.MySQLError)
	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError(config.MessageDuplicateUsername)
	}
	return errors.NewInternalServerError(config.MessageErr)
}
