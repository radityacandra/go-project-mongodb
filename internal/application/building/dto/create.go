package dto

type CreateRequest struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"lat" validate:"required"`
	Longitude float64 `json:"long" validate:"required"`
}

type CreateResponse struct {
	BuildingId string  `json:"buildingId"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"long"`
}

type CountDto struct {
	Total int `bson:"total"`
}
