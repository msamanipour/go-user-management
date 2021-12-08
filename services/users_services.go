package services

import (
	"go-apk-users/app/config"
	"go-apk-users/domain/users"
	"go-apk-users/utils/crypto_utils"
	"go-apk-users/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
	UserInfo     *users.User
)

type usersService struct{}

type usersServiceInterface interface {
	GetLogin(string, string) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	GetUsers() ([]users.User, *errors.RestErr)
	EditUser(user users.User) *errors.RestErr
	DeleteUser(userId int64) *errors.RestErr
	ProfileEdit(userId int64, password string, rePassword string) *errors.RestErr
}

func (s *usersService) GetLogin(username string, password string) (*users.User, *errors.RestErr) {
	result := &users.User{
		Username: username,
		Password: crypto_utils.GetMd5(password),
	}
	if err := result.Validate(false); err != nil {
		return nil, err
	}
	if err := result.Login(); err != nil {
		return nil, err
	}
	UserInfo = result
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(false); err != nil {
		return nil, err
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	saveRes := user.Save()
	return &user, saveRes
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) GetUsers() ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.AllUsers()
}

func (s *usersService) EditUser(user users.User) *errors.RestErr {
	if user.Password != "" {
		user.Password = crypto_utils.GetMd5(user.Password)
	}
	if err := user.Validate(true); err != nil {
		return err
	}

	if err := user.Update(false); err != nil {
		return err
	}
	return nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	result := &users.User{Id: userId}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) ProfileEdit(userId int64, password string, rePassword string) *errors.RestErr {
	if password == rePassword {
		result := &users.User{
			Id:       userId,
			Password: crypto_utils.GetMd5(password),
		}
		err := result.Update(true)
		return err
	} else {
		return errors.NewBadRequestError(config.MessageErrorPassword)
	}

}
