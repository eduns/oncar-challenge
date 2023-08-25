package dtos

type CreateLeadInputDto struct {
	Name            string
	Email           string
	Phone           string
	ChosenVehicleId string
}

type CreateLeadOutputDto struct {
	Id              string
	Name            string
	Email           string
	Phone           string
	ChosenVehicleId string
}
