package users_controllers

import (
	"github.com/gin-gonic/gin"
	"go-apk-users/app/config"
	"go-apk-users/domain/users"
	"go-apk-users/services"
	"go-apk-users/utils/logger"
	"go-apk-users/utils/session_utils"
	"net/http"
)

type User struct {
	name   string
	family string
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
		session_utils.SetSession(c.Writer, c.Request, config.LoginSessionName, "/")
	}
	c.HTML(http.StatusOK, "pages/login", gin.H{
		"title":  config.TitleLogin,
		"result": result,
		"err":    err,
	})
}
func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, "")
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
	//result , err := services.UsersService.GetPermissions()
	c.HTML(http.StatusOK, "pages/users/new", gin.H{
		"title": config.TitleAddUser,
	})
}
func NewPost(c *gin.Context) {
	var user users.User
	if err := c.ShouldBind(&user); err != nil {
		logger.Error(config.MessageBadFormRequest, err)
		return
	}
	_, returnErr := services.UsersService.CreateUser(user)
	c.HTML(http.StatusOK, "pages/users/new", gin.H{
		"title":  config.TitleLogin,
		"result": "",
		"err":    returnErr,
	})
}
