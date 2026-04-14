package cases_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/cases_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/cases_request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Cases_Controller struct {
	service cases_service_interface.Cases_Service_Interface
}

func NewCasesControllerRegistry(cases_service cases_service_interface.Cases_Service_Interface) *Cases_Controller {
	return &Cases_Controller{
		service: cases_service,
	}
}

// @Summary Create a new case
// @Description Create a new case with title, description, incident date, and location
// @Tags Cases
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param case_title formData string true "Case Title"
// @Param case_description formData string true "Case Description"
// @Param incident_date formData string true "Incident Date (YYYY-MM-DD)"
// @Param location formData string true "Case Location"
// @Success 201 {object} map[string]interface{} "Success Create Case"
// @Failure 404 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases/create [post]
func (c *Cases_Controller) Create(ctx *gin.Context) {

	request := new(cases_request.Cases_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(404, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", *request.Incident_Date)

	input := &cases_request.Cases_Dto{
		Case_Title:       request.Case_Title,
		Case_Description: request.Case_Description,
		Incident_Date:    dateParse,
		Location:         request.Location,
		Status:           request.Status,
		DetectiveIDs:     request.DetectiveIDs,
	}

	cases, errCreate := c.service.Create(input)

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
		"Message": "Success Create Case",
		"Data":    cases,
	})
}

// @Summary Get all cases
// @Description Retrieve all cases from the database
// @Tags Cases
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "Success Get Cases"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases [get]
func (c *Cases_Controller) GetAll(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	cases, meta, errGet := c.service.GetAll(page, limit)

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
		"Message":    "Success Get Cases",
		"Pagination": meta,
		"Data":       cases,
	})

}

// @Summary Get case by ID
// @Description Retrieve a specific case using its UUID
// @Tags Cases
// @Produce json
// @Param id path string true "Case ID (UUID)"
// @Success 200 {object} map[string]interface{} "Success Get Case By Id"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases/{id} [get]
func (c *Cases_Controller) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	cases, errGet := c.service.GetById(id)

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
		"Message": "Success Get Case By Id",
		"Data":    cases,
	})

}

// @Summary Get Count Case
// @Description Retrieve a case using its case number
// @Tags Cases
// @Produce json
// @Param number path string true "Case Number"
// @Success 200 {object} map[string]interface{} "Success Get Count Cases"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases/count [get]
func (c *Cases_Controller) GetCount(ctx *gin.Context) {

	CasesCount, errGet := c.service.GetCount()

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
		"Message": "Success Count Cases",
		"Data":    CasesCount,
	})

}

// @Summary Update a case
// @Description Update an existing case with new title, description, incident date, and location
// @Tags Cases
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Case ID (UUID)"
// @Param case_title formData string true "Case Title"
// @Param case_description formData string true "Case Description"
// @Param incident_date formData string true "Incident Date (YYYY-MM-DD)"
// @Param location formData string true "Case Location"
// @Success 200 {object} map[string]interface{} "Success Update Case"
// @Failure 404 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases/update/{id} [put]
func (c *Cases_Controller) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	request := new(cases_request.Cases_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(404, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	dateParse, _ := time.Parse("2006-01-02", *request.Incident_Date)

	input := &cases_request.Cases_Dto{
		Case_Title:       request.Case_Title,
		Case_Description: request.Case_Description,
		Incident_Date:    dateParse,
		Location:         request.Location,
		Status:           request.Status,
		DetectiveIDs:     request.DetectiveIDs,
	}

	cases, errUpdate := c.service.Update(id, input)

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
		"Message": "Success Update Case",
		"Data":    cases,
	})

}

// @Summary Delete a case
// @Description Delete a case by its ID
// @Tags Cases
// @Produce json
// @Param id path string true "Case ID (UUID)"
// @Success 200 {object} map[string]interface{} "Success Delete Case"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/cases/delete/{id} [delete]
func (c *Cases_Controller) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	errDelete := c.service.Delete(id)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Succes Delete Case",
	})

}
