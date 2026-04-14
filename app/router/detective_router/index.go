package detective_router

import (
	"databse-cluster-master-slave-architecture-golang/app/controller/detective_controller"

	"github.com/gin-gonic/gin"
)

func DetectiveRouter(app *gin.Engine, DetectiveController *detective_controller.Detective_Controller) {
	detective := app.Group("/api/detective")

	detective.POST("/create", DetectiveController.Create)
	detective.GET("/", DetectiveController.GetAll)
	detective.GET("/:id", DetectiveController.GetById)
	detective.PUT("/update/:id", DetectiveController.Update)
	detective.DELETE("/delete/:id", DetectiveController.Delete)
}
