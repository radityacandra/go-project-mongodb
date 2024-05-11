package dto

type UpdateRequest struct {
	Id         string  `param:"id" validate:"required"`
	Name       *string `json:"name"`
	BuildingId *string `json:"buildingId"`
	Status     *bool   `json:"status"`
}

type UpdateResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Status     bool   `json:"status"`
	BuildingId string `json:"buildingId"`
}
