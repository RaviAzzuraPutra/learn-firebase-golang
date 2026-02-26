package controller

import (
	"coba-firebase-golang/app/helper"
	"coba-firebase-golang/app/interface/service/member_service_interface"
	"coba-firebase-golang/app/request"

	"github.com/gin-gonic/gin"
)

type Member_Controller struct {
	service member_service_interface.Member_Service_Interface
}

func NewMemberControllerRegistry(Member_Service member_service_interface.Member_Service_Interface) *Member_Controller {
	return &Member_Controller{
		service: Member_Service,
	}
}

func (c *Member_Controller) Create(ctx *gin.Context) {

	request := new(request.Member_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	member, errCreate := c.service.Create(request)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": "Something Wrong",
				"Error":   appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Something Wrong",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Message": "Success",
		"Data":    member,
	})

}

func (c *Member_Controller) GetAll(ctx *gin.Context) {

	member, errGet := c.service.GetAll()

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": "Something Wrong",
				"Error":   appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Something Wrong",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"Data":    member,
	})

}

func (c *Member_Controller) GetById(ctx *gin.Context) {

	ID := ctx.Param("id")

	member, errGet := c.service.GetById(ID)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": "Something Wrong",
				"Error":   appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Something Wrong",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"Data":    member,
	})

}

func (c *Member_Controller) Update(ctx *gin.Context) {

	ID := ctx.Param("id")

	request := new(request.Member_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	member, errUpdate := c.service.Update(ID, request)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": "Something Wrong",
				"Error":   appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Something Wrong",
			"Error":   errUpdate.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
		"Data":    member,
	})

}

func (c *Member_Controller) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": "Something Wrong",
				"Error":   appError.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Something Wrong",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success",
	})

}
