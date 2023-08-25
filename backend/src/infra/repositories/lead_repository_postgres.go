package repositories

import (
	"database/sql"
	"log"

	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
)

type LeadRepositoryPostgres struct {
	connection *sql.DB
}

func NewLeadRepositoryPostgres(conn *sql.DB) LeadRepositoryPostgres {
	return LeadRepositoryPostgres{connection: conn}
}

func (r *LeadRepositoryPostgres) Create(lead entities.Lead) error {
	_, err := r.connection.Exec("INSERT INTO leads (id, name, email, phone, chosenVehicleId) VALUES ($1, $2, $3, $4, $5)",
		lead.Id, lead.Name, lead.Email, lead.Phone, lead.ChosenVehicleId)

	if err != nil {
		log.Println("[DATABASE ERROR] |>", err.Error())
		return err
	}

	return nil
}
