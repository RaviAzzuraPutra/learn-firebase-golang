package bootstrap

import (
	"coba-firebase-golang/app/database/connect"

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

	app.Run(":1221")
}
