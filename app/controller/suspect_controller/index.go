package suspect_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/suspect_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/suspects_request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Suspect_Controller struct {
	service suspect_service_interface.Suspect_Service_Interface
}

func NewSuspectControllerRegistry(suspect_service suspect_service_interface.Suspect_Service_Interface) *Suspect_Controller {
	return &Suspect_Controller{
		service: suspect_service,
	}
}

// @Summary Create a new suspect
// @Description Create a new suspect for a specific case
// @Tags Suspects
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id_case path string true "Case ID (UUID)"
// @Param id_card_number formData string true "ID Card Number"
// @Param full_name formData string true "Full Name"
// @Param address formData string true "Address"
// @Param alibi formData string true "Alibi"
// @Success 201 {object} map[string]interface{} "Success Create Suspect"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/suspect/create/{id_case} [post]
func (c *Suspect_Controller) Create(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: request.ID_Card_Number,
		Full_Name:      request.Full_Name,
		Address:        request.Address,
		Alibi:          request.Alibi,
	}

	suspect, errCreate := c.service.Create(ID_Case, input)

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
		"Messaage": "Success Create Suspect",
		"Data":     suspect,
	})

}

// @Summary Get all suspects
// @Description Retrieve all suspects for a specific case
// @Tags Suspects
// @Produce json
// @Param id_case path string true "Case ID (UUID)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "Success Get Suspect Data"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/suspect/get-all/{id_case} [get]
func (c *Suspect_Controller) GetAll(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	suspect, meta, errGet := c.service.GetAll(ID_Case, page, limit)

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
		"Message":    "Success Get Suspect Data",
		"Data":       suspect,
		"Pagination": meta,
	})

}

// @Summary Get suspect by ID
// @Description Retrieve a specific suspect from a case
// @Tags Suspects
// @Produce json
// @Param id_case path string true "Case ID (UUID)"
// @Param id path string true "Suspect ID (UUID)"
// @Success 200 {object} map[string]interface{} "Success Get Suspect Data"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/suspect/get-id/{id}/{id_case} [get]
func (c *Suspect_Controller) GetById(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	suspect, errGet := c.service.GetById(ID, ID_Case)

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
		"Message": "Success Get Suspect Data",
		"Data":    suspect,
	})

}

// @Summary Update suspect data
// @Description Update an existing suspect information
// @Tags Suspects
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id_case path string true "Case ID (UUID)"
// @Param id path string true "Suspect ID (UUID)"
// @Param id_card_number formData string true "ID Card Number"
// @Param full_name formData string true "Full Name"
// @Param address formData string true "Address"
// @Param alibi formData string true "Alibi"
// @Success 200 {object} map[string]interface{} "Success Update Suspect Data"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/suspect/update/{id}/{id_case} [put]
func (c *Suspect_Controller) Update(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: request.ID_Card_Number,
		Full_Name:      request.Full_Name,
		Address:        request.Address,
		Alibi:          request.Alibi,
	}

	suspect, errUpdate := c.service.Update(ID, ID_Case, input)

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
		"Message": "Success Update Suspect Data",
		"Data":    suspect,
	})

}

// @Summary Delete suspect
// @Description Delete a suspect from a case
// @Tags Suspects
// @Produce json
// @Param id_case path string true "Case ID (UUID)"
// @Param id path string true "Suspect ID (UUID)"
// @Success 200 {object} map[string]interface{} "Success Delete Suspect Data"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/suspect/delete/{id}/{id_case} [delete]
func (c *Suspect_Controller) Delete(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID, ID_Case)

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
		"Message": "Success Delete Suspect Data",
	})

}
