package app

import (
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
	"go-apk-users/domain/users"
	"go-apk-users/services"
	"go-apk-users/utils/authority_utils"
	"html/template"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Static("/static", "./resources/assets")
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "resources/views",
		Extension:    ".html",
		Master:       "layouts/base",
		DisableCache: true,
		Funcs: template.FuncMap{
			"GAuth": func() *authority.Authority {
				return authority_utils.Auth
			},
			"GUser": func() *users.User {
				if services.UserInfo != nil {
					return services.UserInfo
				}
				return &users.User{
					Id:     1,
					Name:   "",
					Family: "",
				}
			},
			"GUint": func(val int64) uint {
				return uint(val)
			},
			"GContain": func(s []string, str string) bool {
				for _, v := range s {
					if v == str {
						return true
					}
				}
				return false
			},
		},
	})
	mapUrls()
	log.Fatal(router.Run(":8080"))
}
