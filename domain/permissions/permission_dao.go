package permissions

import (
	"go-apk-users/app/config"
	"go-apk-users/datasources/mysql/users_db"
	"go-apk-users/utils/errors"
)

func (p *Permission) FetchAll() ([]Permission, *errors.RestErr) {
	var permissions []Permission
	result := users_db.Client.Order("id desc").Find(&permissions)
	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, errors.NewNotFoundError(config.MessageRecordNotFound)
	}
	return permissions, nil
}
