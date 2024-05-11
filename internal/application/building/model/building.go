package model

import "github.com/google/uuid"

type Building struct {
	Id        string  `bson:"id"`
	UserId    string  `bson:"userId"`
	Name      string  `bson:"name"`
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

func NewBuilding(userId, name string, lat, long float64) *Building {
	return &Building{
		Id:        uuid.NewString(),
		UserId:    userId,
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	}
}
