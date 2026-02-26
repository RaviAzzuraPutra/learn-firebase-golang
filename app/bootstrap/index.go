package bootstrap

import (
	"coba-firebase-golang/app/controller"
	"coba-firebase-golang/app/database/connect"
	"coba-firebase-golang/app/repository"
	"coba-firebase-golang/app/router"
	"coba-firebase-golang/app/service"

	"github.com/gin-gonic/gin"
)

func InitApp() {

	connect.Connect()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "The application is running well. 🚀",
		})
	})

	repository := repository.NewMemberRepositoryRegistry()
	service := service.NewMemberServiceRegistry(repository)
	controller := controller.NewMemberControllerRegistry(service)

	router.MemberRouter(app, controller)

	app.Run(":1221")
}
