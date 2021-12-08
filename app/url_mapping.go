package app

import (
	"go-apk-users/app/middleware/auth"
	"go-apk-users/controllers/roles_controllers"
	"go-apk-users/controllers/users_controllers"
)

func mapUrls() {
	router.Use(auth.CheckGlobal)
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
		authed.GET("users/new", auth.Permission("user-create"), users_controllers.New)
		authed.POST("users/new", auth.Permission("user-create"), users_controllers.NewPost)
		authed.GET("users/edit/:user_id", auth.Permission("user-edit"), users_controllers.Edit)
		authed.POST("users/edit/:user_id", auth.Permission("user-edit"), users_controllers.EditPost)
		authed.GET("users/delete/:user_id", auth.Permission("user-delete"), users_controllers.Delete)
		authed.GET("users/profile", users_controllers.Profile)
		authed.POST("users/profile", users_controllers.ProfilePost)
		//roles routes
		authed.GET("roles/manage", roles_controllers.Manage)
		authed.GET("roles/edit/:role_name", roles_controllers.Edit)
		authed.POST("roles/edit/:role_name", roles_controllers.EditPost)
		authed.GET("roles/new", roles_controllers.New)
		authed.POST("roles/new", roles_controllers.NewPost)
	}
}
