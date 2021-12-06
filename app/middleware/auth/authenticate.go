package auth

import (
	"github.com/gin-gonic/gin"
	"go-apk-users/app/config"
	"go-apk-users/utils/session_utils"
	"net/http"
)

func Logged(c *gin.Context) {
	if res := session_utils.CheckLogin(c.Request, "session_token"); res == false {
		c.Redirect(http.StatusFound, config.LoginUrl)
	}
	c.Next()
}
func Guest(c *gin.Context) {
	if res := session_utils.CheckLogin(c.Request, "session_token"); res == true {
		c.Redirect(http.StatusFound, config.DashboardUrl)
	}
	c.Next()
}
