package main

import (
	"go-apk-users/app"
	"go-apk-users/domain/users"
	"go-apk-users/utils/authority_utils"
)

func main() {
	authority_utils.Implement()
	users.Init()
	app.StartApplication()
}
