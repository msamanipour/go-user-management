package users_controllers

import (
	"github.com/gin-gonic/gin"
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
		"title":  config.TitleAddUser,
		"result": user,
		"err":    err,
	})

}

func EditPost(c *gin.Context) {
	//var user users.User
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	restErr := errors.NewBadRequestError("invalid json body")
	//	c.JSON(restErr.Status, restErr.Message)
	//	return
	//}
	//isPartial := c.Request.Method == http.MethodPatch
	//user.Id = userId
	//result, err := services.UsersService.UpdateUser(isPartial, user)
	//if err != nil {
	//	c.JSON(err.Status, err)
	//	return
	//}
	//c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}
