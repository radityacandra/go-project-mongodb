package model

import "github.com/google/uuid"

type Session struct {
	Id        string `bson:"id"`
	ExpiredAt int64  `bson:"expiredAt"`
	IsActive  bool   `bson:"isActive"`
}

type User struct {
	Id          string    `bson:"id"`
	Username    string    `bson:"username"`
	PhoneNumber string    `bson:"phoneNumber"`
	Password    string    `bson:"password"`
	Sessions    []Session `bson:"sessions"`
}

func NewUser(username, phoneNumber, password string) *User {
	return &User{
		Id:          uuid.NewString(),
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
}

func (u *User) AddSession(sessionId string, expiredAt int64) {
	u.Sessions = append(u.Sessions, Session{
		Id:        sessionId,
		ExpiredAt: expiredAt,
		IsActive:  true,
	})
}

func (u *User) GetLastSession() *Session {
	if len(u.Sessions) > 0 {
		return &u.Sessions[(len(u.Sessions) - 1)]
	}

	return nil
}
