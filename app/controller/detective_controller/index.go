package detective_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/detective_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/detective_request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Detective_Controller struct {
	service detective_service_interface.Detective_Service_Interface
}

func NewDetectiveControllerRegistry(detective_service detective_service_interface.Detective_Service_Interface) *Detective_Controller {
	return &Detective_Controller{
		service: detective_service,
	}
}

func (c *Detective_Controller) Create(ctx *gin.Context) {

	request := new(detective_request.Detective_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &detective_request.Detective_Dto{
		Name:                request.Name,
		Badge_Number:        request.Badge_Number,
		Department:          request.Department,
		Station:             request.Station,
		Phone:               request.Phone,
		Investigation_Style: request.Investigation_Style,
	}

	detective, errCreate := c.service.Create(input)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errCreate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Create Detective",
		"Data":    detective,
	})

}

func (c *Detective_Controller) GetAll(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	detective, meta, errGet := c.service.GetAll(page, limit)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message":    "Success Get Detective",
		"Pagination": meta,
		"Data":       detective,
	})

}

func (c *Detective_Controller) GetById(ctx *gin.Context) {

	ID := ctx.Param("id")

	detective, errGet := c.service.GetById(ID)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Data By Id",
		"Data":    detective,
	})
}

func (c *Detective_Controller) Update(ctx *gin.Context) {

	ID := ctx.Param("id")

	request := new(detective_request.Detective_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &detective_request.Detective_Dto{
		Name:                request.Name,
		Badge_Number:        request.Badge_Number,
		Department:          request.Department,
		Station:             request.Station,
		Phone:               request.Phone,
		Investigation_Style: request.Investigation_Style,
	}

	detective, errUpdate := c.service.Update(ID, input)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errUpdate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errUpdate.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Update Detective",
		"Data":    detective,
	})

}

func (c *Detective_Controller) Delete(ctx *gin.Context) {

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Delete Detective",
	})
}
