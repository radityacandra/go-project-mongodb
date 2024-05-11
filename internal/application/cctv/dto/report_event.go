package dto

type ReportEventRequest struct {
	Id        string  `param:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Timestamp int64   `json:"timestamp" validate:"required"`
	Type      string  `json:"type" validate:"required,oneof=info warning alarm"`
	Source    *string `json:"source"` // Need some mechanism on this, but let's assume just so
}

type ReportEventResponse struct {
	EventId   string  `json:"eventId"`
	Name      string  `json:"name"`
	Timestamp int64   `json:"timestamp"`
	Type      string  `json:"type"`
	Source    *string `json:"source"` // Need some mechanism on this, but let's assume just so
}
