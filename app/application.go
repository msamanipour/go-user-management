package app

import (
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
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
	})
	mapUrls()
	log.Fatal(router.Run(":8080"))
}
