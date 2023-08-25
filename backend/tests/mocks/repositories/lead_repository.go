package repositories

import (
	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
)

type LeadRepositoryInMemory struct {
	leads []entities.Lead
}

func NewLeadRepositoryInMemory() LeadRepositoryInMemory {
	return LeadRepositoryInMemory{}
}

func (r *LeadRepositoryInMemory) Create(lead entities.Lead) error {
	r.leads = append(r.leads, lead)

	return nil
}
