package model

import "github.com/google/uuid"

type User struct {
	Id          string `bson:"id"`
	Username    string `bson:"username"`
	PhoneNumber string `bson:"phoneNumber"`
	Password    string `bson:"password"`
}

func NewUser(username, phoneNumber, password string) *User {
	return &User{
		Id:          uuid.NewString(),
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
}
