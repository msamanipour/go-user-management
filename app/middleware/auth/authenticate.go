package auth

import (
	"github.com/gin-gonic/gin"
	"go-apk-users/app/config"
	"go-apk-users/services"
	"go-apk-users/utils/session_utils"
	"net/http"
)

func Logged(c *gin.Context) {
	if res := session_utils.CheckSession(c.Request, "session_token"); res == false {
		c.Redirect(http.StatusFound, config.LoginUrl)
	}
	c.Next()
}
func Guest(c *gin.Context) {
	if res := session_utils.CheckSession(c.Request, "session_token"); res == true {
		c.Redirect(http.StatusFound, config.DashboardUrl)
	}
	c.Next()
}

func CheckGlobal(c *gin.Context) {
	if services.UserInfo == nil {
		session_utils.ClearSession(c.Writer, config.LoginSessionName)
		if c.FullPath() != config.LoginUrl {
			c.Redirect(http.StatusFound, config.LoginUrl)
		}
	}
}
