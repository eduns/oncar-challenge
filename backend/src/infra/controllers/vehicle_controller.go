package controllers

import (
	"net/http"

	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/src/interfaces"
	"github.com/gin-gonic/gin"
)

type VehiclesController struct {
	getVehiclesUseCase usecases.GetVehiclesUseCase
}

func NewVehicleController(usecase usecases.GetVehiclesUseCase) VehiclesController {
	return VehiclesController{getVehiclesUseCase: usecase}
}

func (c *VehiclesController) HandleGetVehicles(ctx *gin.Context) {
	vehicles, err := c.getVehiclesUseCase.Execute()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, vehicles)
}
