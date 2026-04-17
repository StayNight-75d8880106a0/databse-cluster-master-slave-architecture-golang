package cases_router

import (
	"databse-cluster-master-slave-architecture-golang/app/controller/cases_controller"

	"github.com/gin-gonic/gin"
)

func CasesRouter(app *gin.Engine, CasesController *cases_controller.Cases_Controller) {

	cases := app.Group("/api/cases")

	cases.POST("/create", CasesController.Create)
	cases.GET("/", CasesController.GetAll)
	cases.GET("/latest", CasesController.GetCasesLatestUpdate)
	cases.GET("/:id", CasesController.GetById)
	cases.GET("/count", CasesController.GetCount)
	cases.PUT("/update/:id", CasesController.Update)
	cases.DELETE("/delete/:id", CasesController.Delete)

}
