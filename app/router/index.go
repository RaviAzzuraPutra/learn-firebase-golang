package router

import (
	"coba-firebase-golang/app/controller"

	"github.com/gin-gonic/gin"
)

func MemberRouter(app *gin.Engine, MemberController *controller.Member_Controller) {

	member := app.Group("/api/member")

	member.POST("/create", MemberController.Create)
	member.GET("/", MemberController.GetAll)
	member.GET("/:id", MemberController.GetById)
	member.PUT("/update/:id", MemberController.Update)
	member.DELETE("/delete/:id", MemberController.Delete)

}
