package dtos

type GetVehiclesOutputDto struct {
	Id    string  `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  uint16  `json:"year"`
	Price float32 `json:"price"`
}
