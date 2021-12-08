package roles_controllers

import (
	"github.com/gin-gonic/gin"
	"go-apk-users/app/config"
	"go-apk-users/utils/authority_utils"
	"go-apk-users/utils/errors"
	"net/http"
)

func Manage(c *gin.Context) {
	roles, err := authority_utils.Auth.GetRoles()
	permissions, err := authority_utils.Auth.GetPermissions()
	c.HTML(http.StatusOK, "pages/roles/manager", gin.H{
		"title":       config.TitlePermissionManager,
		"roles":       roles,
		"permissions": permissions,
		"err":         err,
	})
}

func Edit(c *gin.Context) {
	allPermissions, _ := authority_utils.Auth.GetPermissions()
	permissions := authority_utils.Auth.GetRolePermissions(c.Param("role_name"))
	c.HTML(http.StatusOK, "pages/roles/editRole", gin.H{
		"title":          config.TitleEditRole,
		"allPermissions": allPermissions,
		"permissions":    permissions,
		"role_name":      c.Param("role_name"),
	})
}

func EditPost(c *gin.Context) {
	_ = authority_utils.Auth.AssignPermissions(c.PostForm("role_name"), c.PostFormArray("permissions"))
	allPermissions, _ := authority_utils.Auth.GetPermissions()
	permissions := authority_utils.Auth.GetRolePermissions(c.Param("role_name"))
	c.HTML(http.StatusOK, "pages/roles/editRole", gin.H{
		"title":          config.TitleEditRole,
		"allPermissions": allPermissions,
		"permissions":    permissions,
		"err":            errors.NewSuccessMessage(config.MessageSuccessEditRole),
	})
}

func New(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/roles/newRole", gin.H{
		"title": config.TitleEditRole,
	})
}

func NewPost(c *gin.Context) {
	err := authority_utils.Auth.CreateRole(c.PostForm("name"))
	var reErr *errors.RestErr
	if err != nil {
		reErr = errors.NewInternalServerError(config.MessageErr)
	} else {
		reErr = errors.NewSuccessMessage(config.MessageSuccessAddRole)
	}
	c.HTML(http.StatusOK, "pages/roles/newRole", gin.H{
		"title": config.TitleEditRole,
		"err":   reErr,
	})
}
