package usecases

import (
	"testing"

	"github.com/eduns/oncar-challenge/backend/src/domain/dtos"
	"github.com/eduns/oncar-challenge/backend/src/domain/usecases"
	"github.com/eduns/oncar-challenge/backend/tests/mocks/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateLeadUseCase(t *testing.T) {
	leadRepository := repositories.NewLeadRepositoryInMemory()
	uc := usecases.NewCreateLeadUseCase(&leadRepository)

	input := dtos.CreateLeadInputDto{
		Name:            "Test",
		Email:           "mail@mail.com",
		Phone:           "+5512881437331",
		ChosenVehicleId: "1da55e15-f253-488d-8590-e0c8a955898b",
	}

	createdLead, err := uc.Execute(input)

	assert.Nil(t, err, "Should not return error")

	assert.NotNil(t, createdLead.Id, "Lead id should not be null")
	assert.Equal(t, createdLead.Name, "Test", "Lead name should not be null")
	assert.Equal(t, createdLead.Email, "mail@mail.com", "Lead email should be present")
	assert.Equal(t, createdLead.Phone, "+5512881437331", "Lead phone should be present")
	assert.Equal(t, createdLead.ChosenVehicleId, "1da55e15-f253-488d-8590-e0c8a955898b", "Lead chosenVehicleId should not be null")
}
