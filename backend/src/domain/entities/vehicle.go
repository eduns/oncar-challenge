package entities

import "github.com/google/uuid"

type Vehicle struct {
	Id    string
	Brand string
	Model string
	Year  uint16
	Price float32
}

func NewVehicle(brand string, model string, year uint16, price float32) Vehicle {
	return Vehicle{
		Id:    uuid.New().String(),
		Brand: brand,
		Model: model,
		Year:  year,
		Price: price,
	}
}
