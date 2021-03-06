package users

import (
	"go-apk-users/app/config"
	"go-apk-users/datasources/mysql/users_db"
	"go-apk-users/utils/authority_utils"
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
	result := users_db.Client.Select("id", "name", "family", "username", "created_at", "updated_at", "password").Create(&u)
	err := authority_utils.Auth.AssignRole(uint(u.Id), u.Role)
	if result.Error != nil {
		logger.Error("error in user save", result.Error)
		return mysql_utils.ParseErrors(result.Error)
	}
	if err != nil {
		logger.Error("error in assign role", err)
		return errors.NewInternalServerError(config.MessageErr)
	}
	return errors.NewSuccessMessage(config.MessageSuccessCreateUser)
}

func (u *User) Get() *errors.RestErr {
	result := users_db.Client.Find(&u)
	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.NewNotFoundError(config.MessageUserNotFound)
	}
	return nil
}

func (u *User) AllUsers() ([]User, *errors.RestErr) {
	var users []User
	result := users_db.Client.Where("id != ?", "1").Order("id desc").Find(&users)
	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, errors.NewNotFoundError(config.MessageUserNotFound)
	}
	return users, nil
}

func (u *User) Update(isProfile bool) *errors.RestErr {
	if !isProfile {
		result := users_db.Client.Select("name", "family", "username", "password").Updates(&u)
		if result.Error != nil {
			return errors.NewNotFoundError(config.MessageErr)
		}
		oldRole, err := authority_utils.Auth.GetUserRoles(uint(u.Id))
		if err != nil {
			return nil
		}
		reErr := authority_utils.Auth.RevokeRole(uint(u.Id), oldRole[0])
		if reErr != nil {
			logger.Error("error in revoke role", reErr)
			return errors.NewInternalServerError(config.MessageErr)
		}
		sErr := authority_utils.Auth.AssignRole(uint(u.Id), u.Role)
		if sErr != nil {
			logger.Error("error in assign role", sErr)
			return errors.NewInternalServerError(config.MessageErr)
		}
	} else {
		result := users_db.Client.Model(&u).Where("id = ?", u.Id).Update("password", u.Password)
		if result.Error != nil {
			logger.Error("error in edit profile", result.Error)
			return errors.NewNotFoundError(config.MessageErr)
		}
	}
	return errors.NewSuccessMessage(config.MessageSuccessEditUser)
}

func (u *User) Delete() *errors.RestErr {
	oldRole, err := authority_utils.Auth.GetUserRoles(uint(u.Id))
	if err != nil {
		logger.Error("error in get user role", err)
		return nil
	}
	reErr := authority_utils.Auth.RevokeRole(uint(u.Id), oldRole[0])
	if reErr != nil {
		logger.Error("error in revoke role", reErr)
		return errors.NewInternalServerError(config.MessageErr)
	}
	result := users_db.Client.Delete(&u)
	if result.Error != nil {
		return errors.NewNotFoundError(config.MessageErr)
	}
	return errors.NewSuccessMessage(config.MessageSuccessDeleteUser)
}
