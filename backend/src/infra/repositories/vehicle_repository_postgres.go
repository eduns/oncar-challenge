package repositories

import (
	"database/sql"
	"log"

	"github.com/eduns/oncar-challenge/backend/src/domain/entities"
)

type VehicleRepositoryPostgres struct {
	connection *sql.DB
}

func NewVehicleRepositoryPostgres(conn *sql.DB) VehicleRepositoryPostgres {
	return VehicleRepositoryPostgres{connection: conn}
}

func (r *VehicleRepositoryPostgres) Find() ([]entities.Vehicle, error) {
	results, err := r.connection.Query("SELECT * FROM vehicles")

	if err != nil {
		log.Println("[DATABASE ERROR] |>", err.Error())
		return []entities.Vehicle{}, err
	}

	defer results.Close()

	var vehicles []entities.Vehicle = []entities.Vehicle{}

	for results.Next() {
		var vehicle entities.Vehicle

		if err := results.Scan(&vehicle.Id, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Price); err != nil {
			log.Println("[DATABASE ERROR] |>", err.Error())
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}
