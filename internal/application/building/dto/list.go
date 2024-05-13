package dto

type ListRequest struct {
	Page    int     `query:"page" validate:"omitempty,min=1"`
	PerPage int     `query:"per_page" validate:"omitempty,min=1"`
	Sort    string  `query:"sort" validate:"omitempty,oneof=name lat long"`
	Order   string  `query:"order" validate:"omitempty,oneof=asc desc"`
	Keyword *string `query:"keyword"`
}

type ListResponse struct {
	BuildingId string  `json:"buildingId"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"long"`
}

type MetaResponse struct {
	MaxPage int
	Page    int
	PerPage int
}
