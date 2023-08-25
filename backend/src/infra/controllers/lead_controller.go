package controllers

import (
	"net/http"

	"github.com/eduns/oncar-challenge/backend/src/domain/dtos"
	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/src/interfaces"
	"github.com/gin-gonic/gin"
)

type CreateLeadPayload struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Phone           string `json:"phone" binding:"required,e164"`
	ChosenVehicleId string `json:"chosenVehicleId" binding:"required,uuid4"`
}

type LeadController struct {
	createLeadUseCase usecases.CreateLeadUseCase
}

func NewLeadController(usecase usecases.CreateLeadUseCase) LeadController {
	return LeadController{createLeadUseCase: usecase}
}

func (c *LeadController) HandleCreateLead(ctx *gin.Context) {
	var payload CreateLeadPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	_, err := c.createLeadUseCase.Execute(dtos.CreateLeadInputDto{
		Name:            payload.Name,
		Email:           payload.Email,
		Phone:           payload.Phone,
		ChosenVehicleId: payload.ChosenVehicleId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Error: err.Error(),
		})

		return
	}

	ctx.Status(http.StatusCreated)
}
