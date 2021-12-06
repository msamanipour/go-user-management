package main

import (
	"go-apk-users/app"
	"go-apk-users/utils/authority_utils"
)

func main() {
	authority_utils.Implement()
	app.StartApplication()
}
