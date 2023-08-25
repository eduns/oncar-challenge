package usecases

import (
	"github.com/eduns/oncar-challenge/backend/src/domain/dtos"
	"github.com/eduns/oncar-challenge/backend/src/interfaces"
)

type GetVehiclesUseCase struct {
	vehicleRepository interfaces.VehicleRepository
}

func NewGetVehiclesUseCase(repository interfaces.VehicleRepository) GetVehiclesUseCase {
	return GetVehiclesUseCase{vehicleRepository: repository}
}

func (u *GetVehiclesUseCase) Execute() ([]dtos.GetVehiclesOutputDto, error) {
	results, err := u.vehicleRepository.Find()

	if err != nil {
		return nil, err
	}

	var vehicles []dtos.GetVehiclesOutputDto = []dtos.GetVehiclesOutputDto{}

	for _, result := range results {
		vehicles = append(vehicles, dtos.GetVehiclesOutputDto{
			Id:    result.Id,
			Brand: result.Brand,
			Model: result.Model,
			Year:  result.Year,
			Price: result.Price,
		})
	}

	return vehicles, nil
}
