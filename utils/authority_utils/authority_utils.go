package authority_utils

import (
	"github.com/harranali/authority"
	"go-apk-users/datasources/mysql/users_db"
)

var (
	Auth *authority.Authority
)

func Implement() {
	Auth = authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           users_db.Client,
	})
	_ = Auth.CreateRole("admin")
	_ = Auth.CreateRole("users-manager")
	_ = Auth.CreateRole("users-supervisor")

	_ = Auth.CreatePermission("user-create")
	_ = Auth.CreatePermission("user-delete")
	_ = Auth.CreatePermission("user-edit")
	_ = Auth.CreatePermission("role-manager")

	_ = Auth.AssignPermissions("admin", []string{
		"user-create",
		"user-delete",
		"user-edit",
		"role-manager",
	})
	_ = Auth.AssignPermissions("users-manager", []string{
		"user-create",
		"user-edit",
	})
	_ = Auth.AssignPermissions("users-supervisor", []string{
		"user-create",
	})
}
