package users_controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-apk-users/app/config"
	"go-apk-users/domain/users"
	"go-apk-users/services"
	"go-apk-users/utils/authority_utils"
	"go-apk-users/utils/errors"
	"go-apk-users/utils/logger"
	"go-apk-users/utils/session_utils"
	"net/http"
	"strconv"
)

func getUserId(uIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(uIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Index(c *gin.Context) {
	c.Redirect(http.StatusFound, config.LoginUrl)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/login", gin.H{
		"title": config.TitleLogin,
	})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result, err := services.UsersService.GetLogin(username, password)
	if err == nil {
		session_utils.SetSession(c.Writer, c.Request, config.LoginSessionName, "/", uuid.NewV4().String())
		//session_utils.SetSession(c.Writer, c.Request, "session_info", "/", crypto_utils.Base64Encode(strconv.FormatInt(result.Id, 10)))
		c.Redirect(http.StatusFound, config.DashboardUrl)
	} else {
		c.HTML(http.StatusOK, "pages/login", gin.H{
			"title":  config.TitleLogin,
			"result": result,
			"err":    err,
		})
	}
}

func Logout(c *gin.Context) {
	session_utils.ClearSession(c.Writer, config.LoginSessionName)
	c.Redirect(http.StatusFound, config.LoginUrl)
}

func Dashboard(c *gin.Context) {
	result, err := services.UsersService.GetUsers()
	c.HTML(http.StatusOK, "pages/dashboard", gin.H{
		"title":  config.TitleDashboard,
		"result": result,
		"err":    err,
	})
}

func New(c *gin.Context) {
	roles, err := authority_utils.Auth.GetRoles()
	c.HTML(http.StatusOK, "pages/users/new", gin.H{
		"title":  config.TitleAddUser,
		"result": roles,
		"err":    err,
	})
}

func NewPost(c *gin.Context) {
	var user users.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Error(config.MessageBadFormRequest, err)
		return
	}
	_, returnErr := services.UsersService.CreateUser(user)
	roles, _ := authority_utils.Auth.GetRoles()
	c.HTML(http.StatusOK, "pages/users/new", gin.H{
		"title":  config.TitleLogin,
		"result": roles,
		"err":    returnErr,
	})
}

func Edit(c *gin.Context) {
	userId, _ := getUserId(c.Param("user_id"))
	user, err := services.UsersService.GetUser(userId)
	c.HTML(http.StatusOK, "pages/users/edit", gin.H{
		"title":   config.TitleAddUser,
		"result":  user,
		"err":     err,
		"user_id": uint(userId),
	})
}

func EditPost(c *gin.Context) {
	var user users.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Error(config.MessageBadFormRequest, err)
		return
	}
	err := services.UsersService.EditUser(user)
	c.HTML(http.StatusOK, "pages/users/edit", gin.H{
		"title":   config.TitleEditUser,
		"result":  user,
		"err":     err,
		"user_id": uint(user.Id),
	})
}

func Delete(c *gin.Context) {
	userId, _ := getUserId(c.Param("user_id"))
	_ = services.UsersService.DeleteUser(userId)
	c.Redirect(http.StatusFound, config.DashboardUrl)
}

func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/users/profile", gin.H{
		"title": config.TitleProfile,
	})
}

func ProfilePost(c *gin.Context) {
	err := services.UsersService.ProfileEdit(services.UserInfo.Id, c.PostForm("password"), c.PostForm("re_password"))
	c.HTML(http.StatusOK, "pages/users/profile", gin.H{
		"title": config.TitleProfile,
		"err":   err,
	})

}
