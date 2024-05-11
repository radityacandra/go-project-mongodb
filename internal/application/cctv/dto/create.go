package dto

type CreateRequest struct {
	Name     string `json:"name" validate:"required"`
	DeviceId string `json:"deviceId" validate:"required"`
}

type CreateResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	DeviceId   string `json:"deviceId"`
	Status     bool   `json:"status"`
	LastUpdate int64  `json:"lastUpdate"`
}
