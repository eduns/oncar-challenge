package interfaces

import (
	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
)

type VehicleRepository interface {
	Find() ([]entities.Vehicle, error)
}
