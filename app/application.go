package app

import (
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
	"go-apk-users/domain/users"
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
		Root:      "resources/views",
		Extension: ".html",
		Master:    "layouts/base",
		//Partials:  []string{"partials/ad"},
		DisableCache: true,
		Funcs: template.FuncMap{
			"GAuth": func() *authority.Authority {
				return authority_utils.Auth
			},
			"GUser": func() *users.User {
				return nil
			},
		},
	})
	mapUrls()
	log.Fatal(router.Run(":8080"))
}
