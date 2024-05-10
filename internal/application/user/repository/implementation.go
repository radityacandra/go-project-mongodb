package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) Save(ctx context.Context, user *model.User) error {
	_, err := r.db.Collection("users").InsertOne(ctx, user)
	return err
}

func (r *Repository) GetByUsername(ctx context.Context, userName string) *model.User {
	var user model.User

	result := r.db.Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: userName}})

	if result.Err() != nil {
		return nil
	}

	if err := result.Decode(&user); err != nil {
		return nil
	}

	return &user
}
