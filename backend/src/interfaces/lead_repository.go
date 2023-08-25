package interfaces

import (
	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
)

type LeadRepository interface {
	Create(lead entities.Lead) error
}
