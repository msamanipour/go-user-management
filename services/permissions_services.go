package services

import (
	"go-apk-users/domain/users"
	"go-apk-users/utils/errors"
)

var (
	PermissionsService permissionsServiceInterface = &permissionsService{}
)

type permissionsService struct{}

type permissionsServiceInterface interface {
	GetPermissions() ([]users.User, *errors.RestErr)
}

func (s *permissionsService) GetPermissions() ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.AllUsers()
}
