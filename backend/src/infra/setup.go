package infra

import (
	"database/sql"

	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/src/infra/controllers"
	"github.com/eduns/oncar-challenge/backend/src/infra/repositories"
	"github.com/gin-gonic/gin"
)

func SetupServer(dbConnection *sql.DB) *gin.Engine {
	vehicleRepository := repositories.NewVehicleRepositoryPostgres(dbConnection)
	leadRepository := repositories.NewLeadRepositoryPostgres(dbConnection)

	getVehiclesUseCase := usecases.NewGetVehiclesUseCase(&vehicleRepository)
	createLeadUseCase := usecases.NewCreateLeadUseCase(&leadRepository)

	vehicleController := controllers.NewVehicleController(getVehiclesUseCase)
	leadController := controllers.NewLeadController(createLeadUseCase)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/leads", leadController.HandleCreateLead)
	router.GET("/vehicles", vehicleController.HandleGetVehicles)

	return router
}
