package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

const COLLECTION_NAME = "users"

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

func (r *Repository) SaveSession(ctx context.Context, user *model.User) error {
	_, err := r.db.Collection(COLLECTION_NAME).UpdateOne(ctx, bson.D{{Key: "id", Value: user.Id}},
		bson.D{{Key: "$push", Value: bson.D{
			{Key: "sessions", Value: user.GetLastSession()},
		}}})

	return err
}

func (r *Repository) GetUserBySessionId(ctx context.Context, userId, sessionId string) *model.User {
	result := r.db.Collection(COLLECTION_NAME).
		FindOne(ctx, bson.D{
			{Key: "id", Value: userId},
			{Key: "sessions", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "id", Value: sessionId}, {Key: "isActive", Value: true}}}}},
		})

	if result.Err() != nil {
		return nil
	}

	var user model.User
	if err := result.Decode(&user); err != nil {
		return nil
	}

	return &user
}

func (r *Repository) RevokeSession(ctx context.Context, userId, sessionId string) error {
	_, err := r.db.Collection(COLLECTION_NAME).
		UpdateOne(ctx,
			bson.D{{Key: "id", Value: userId}, {Key: "sessions.id", Value: sessionId}},
			bson.D{{Key: "$set", Value: bson.D{{Key: "sessions.$.isActive", Value: false}}}})

	return err
}
