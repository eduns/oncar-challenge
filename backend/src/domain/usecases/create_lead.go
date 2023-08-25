package usecases

import (
	"github.com/eduns/oncar-challenge/backend/src/domain/dtos"
	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
	"github.com/eduns/oncar-challenge/backend/src/interfaces"
)

type CreateLeadUseCase struct {
	leadRepository interfaces.LeadRepository
}

func NewCreateLeadUseCase(repository interfaces.LeadRepository) CreateLeadUseCase {
	return CreateLeadUseCase{leadRepository: repository}
}

func (u *CreateLeadUseCase) Execute(input dtos.CreateLeadInputDto) (dtos.CreateLeadOutputDto, error) {
	newLead := entities.NewLead(
		input.Name,
		input.Email,
		input.Phone,
		input.ChosenVehicleId,
	)

	err := u.leadRepository.Create(newLead)

	if err != nil {
		return dtos.CreateLeadOutputDto{}, err
	}

	return dtos.CreateLeadOutputDto{
		Id:              newLead.Id,
		Name:            newLead.Name,
		Email:           newLead.Email,
		Phone:           newLead.Phone,
		ChosenVehicleId: newLead.ChosenVehicleId,
	}, nil
}
