package entities

import "github.com/google/uuid"

type Lead struct {
	Id              string
	Name            string
	Email           string
	Phone           string
	ChosenVehicleId string
}

func NewLead(name string, email string, phone string, chosenVehicleId string) Lead {
	return Lead{
		Id:              uuid.New().String(),
		Name:            name,
		Email:           email,
		Phone:           phone,
		ChosenVehicleId: chosenVehicleId,
	}
}
