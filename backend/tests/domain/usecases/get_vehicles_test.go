package usecases

import (
	"testing"

	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/tests/mocks/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetVehiclesUseCase(t *testing.T) {
	vehicleRepository := repositories.NewVehicleRepositoryInMemory()
	uc := usecases.NewGetVehiclesUseCase(&vehicleRepository)

	result, err := uc.Execute()

	assert.Nil(t, err, "Should not return error")
	assert.Equal(t, 3, len(result), "It should return all vehicles from DB")
}
