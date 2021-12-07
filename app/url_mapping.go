package app

import (
	"go-apk-users/app/middleware/auth"
	"go-apk-users/controllers/users_controllers"
)

func mapUrls() {
	router.GET("/", users_controllers.Index)
	// guest routes
	guested := router.Group("/")
	guested.Use(auth.Guest)
	{
		guested.GET("users/login", users_controllers.Login)
		guested.POST("users/login", users_controllers.LoginPost)
	}
	//auth routes
	authed := router.Group("/")
	authed.Use(auth.Logged)
	{
		authed.GET("/logout", users_controllers.Logout)
		authed.GET("/dashboard", users_controllers.Dashboard)
		authed.GET("users/new", users_controllers.New)
		authed.POST("users/new", users_controllers.NewPost)
		authed.GET("users/edit/:user_id", users_controllers.Edit)
		authed.POST("users/edit/:user_id", users_controllers.EditPost)
	}
}
