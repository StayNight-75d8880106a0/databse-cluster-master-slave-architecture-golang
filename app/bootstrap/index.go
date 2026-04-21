package bootstrap

import (
	"databse-cluster-master-slave-architecture-golang/app/config"
	"databse-cluster-master-slave-architecture-golang/app/config/app_config"
	"databse-cluster-master-slave-architecture-golang/app/controller/ws_controller"
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/registry/ai_registry"
	"databse-cluster-master-slave-architecture-golang/app/registry/cases_registry"
	"databse-cluster-master-slave-architecture-golang/app/registry/detective_registry"
	"databse-cluster-master-slave-architecture-golang/app/registry/suspect_registry"
	"databse-cluster-master-slave-architecture-golang/app/router/cases_router"
	"databse-cluster-master-slave-architecture-golang/app/router/detective_router"
	"databse-cluster-master-slave-architecture-golang/app/router/suspect_router"
	_ "databse-cluster-master-slave-architecture-golang/docs"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAPP() {
	_ = godotenv.Load()

	FrontendURL := os.Getenv("FRONTEND_URL")

	config.Config()
	database.Connect()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"content-length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "The application is running well. 💮",
		})
	})

	AI_Module := ai_registry.AI_Registry(database.GetInstanceDbCluster().Master)

	wsCtrl := ws_controller.NewWebSocketControllerRegistry(AI_Module.AI_Controller)
	app.GET("/ws", wsCtrl.Connect)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	CasesModule := cases_registry.Case_Registry(AI_Module.VectorStore)
	SuspectModule := suspect_registry.Suspect_Registry(AI_Module.VectorStore)
	DetectiveModule := detective_registry.Detective_Registry(AI_Module.VectorStore)

	cases_router.CasesRouter(app, CasesModule.Cases_Controller)
	suspect_router.SuspectRouter(app, SuspectModule.Suspect_Controller)
	detective_router.DetectiveRouter(app, DetectiveModule.Detective_Controller)

	app.Run(":" + app_config.PORT)
}
