package repositories

import "github.com/eduns/oncar-challenge/backend/src/domain/entities"

type VehicleRepositoryInMemory struct {
	vehicles []entities.Vehicle
}

func NewVehicleRepositoryInMemory() VehicleRepositoryInMemory {
	return VehicleRepositoryInMemory{
		vehicles: []entities.Vehicle{
			{
				Id:    "320d3c04-cc57-435e-92f9-d6e3f6e0a7f4",
				Brand: "Tesla",
				Model: "Model S",
				Year:  2022,
				Price: 216000,
			},
			{
				Id:    "1da55e15-f253-488d-8590-e0c8a955898b",
				Brand: "Ford",
				Model: "Mustang",
				Year:  2021,
				Price: 175000,
			},
			{
				Id:    "bc169144-f3bd-440f-8f60-5b9c833eda47",
				Brand: "Ferrari",
				Model: "LaFerrari",
				Year:  2016,
				Price: 572000,
			},
		},
	}
}

func (r *VehicleRepositoryInMemory) Find() ([]entities.Vehicle, error) {
	return r.vehicles, nil
}
