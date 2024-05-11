package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        string  `bson:"id"`
	Name      string  `bson:"name"`
	Timestamp int64   `bson:"timestamp"`
	Type      string  `bson:"type"`
	Source    *string `bson:"source"`
}

type Cctv struct {
	Id         string  `bson:"id"`
	Name       string  `bson:"name"`
	DeviceId   string  `bson:"deviceId"`
	Status     bool    `bson:"status"`
	LastUpdate int64   `bson:"lastUpdate"`
	BuildingId string  `bson:"buildingId"`
	UserId     string  `bson:"userId"`
	Events     []Event `bson:"events"`
}

func NewCctv(userId, name, deviceId string) *Cctv {
	return &Cctv{
		Id:         uuid.NewString(),
		Name:       name,
		DeviceId:   deviceId,
		Status:     true,
		LastUpdate: time.Now().Unix(),
		UserId:     userId,
		Events: []Event{
			{
				Id:        uuid.NewString(),
				Name:      "CCTV Registered",
				Timestamp: time.Now().Unix(),
				Type:      "Info",
			},
		},
	}
}

func (cctv *Cctv) UpdateName(name *string) {
	if name != nil {
		cctv.Name = *name
	}
}

func (cctv *Cctv) UpdateBuildingId(buildingId *string) {
	if buildingId != nil {
		cctv.BuildingId = *buildingId
	}
}

func (cctv *Cctv) UpdateStatus(status *bool) {
	if status != nil {
		cctv.Status = *status
	}

	if status != nil && !*status {
		cctv.Events = append(cctv.Events, Event{
			Id:        uuid.NewString(),
			Name:      "CCTV is shutted down",
			Timestamp: time.Now().Unix(),
			Type:      "warning",
		})
	}
}

func (cctv *Cctv) AddEvent(name string, timestamp int64, eventType string, source *string) Event {
	event := Event{
		Id:        uuid.NewString(),
		Name:      name,
		Timestamp: timestamp,
		Type:      eventType,
		Source:    source,
	}
	cctv.Events = append(cctv.Events, event)

	return event
}
