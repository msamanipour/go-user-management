package services

import (
	"go-apk-users/domain/permissions"
	"go-apk-users/utils/errors"
)

var (
	PermissionsService permissionsServiceInterface = &permissionsService{}
)

type permissionsService struct{}

type permissionsServiceInterface interface {
	GetPermissions() ([]permissions.Permission, *errors.RestErr)
}

func (s *permissionsService) GetPermissions() ([]permissions.Permission, *errors.RestErr) {
	dao := &permissions.Permission{}
	return dao.FetchAll()
}
