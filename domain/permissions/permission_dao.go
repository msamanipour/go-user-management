package users

import (
	"go-apk-users/app/config"
	"go-apk-users/datasources/mysql/users_db"
	"go-apk-users/utils/errors"
	"go-apk-users/utils/logger"
	"go-apk-users/utils/mysql_utils"
)

func (u *User) Login() *errors.RestErr {
	result := users_db.Client.Where("username = ? AND password = ?", u.Username, u.Password).First(&u)
	if result.Error != nil {
		logger.Error("user not found", result.Error)
		return errors.NewNotFoundError(config.MessageWrongUserPass)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	result := users_db.Client.Create(&u)
	if result.Error != nil {
		logger.Error("error in user save", result.Error)
		return mysql_utils.ParseErrors(result.Error)
	}
	return errors.NewSuccessMessage(config.MessageSuccessCreateUser)
}

func (u *User) Get() *errors.RestErr {

	return nil
}

func (u *User) AllUsers() ([]User, *errors.RestErr) {
	var users []User
	result := users_db.Client.Order("id desc").Find(&users)
	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, errors.NewNotFoundError(config.MessageUserNotFound)
	}
	return users, nil
}
