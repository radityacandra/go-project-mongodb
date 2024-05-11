package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/model"
	"go.mongodb.org/mongo-driver/bson"
)

const COLLECTION_NAME = "cctvs"

func (r *Repository) Create(ctx context.Context, cctv *model.Cctv) error {
	_, err := r.db.Collection(COLLECTION_NAME).InsertOne(ctx, cctv)
	return err
}

func (r *Repository) FindByUserIdAndName(ctx context.Context, userId, name string) *model.Cctv {
	return nil
}

func (r *Repository) FindById(ctx context.Context, userId, id string) *model.Cctv {
	result := r.db.Collection(COLLECTION_NAME).FindOne(ctx, bson.D{{Key: "userId", Value: userId}, {Key: "id", Value: id}})

	if result.Err() != nil {
		return nil
	}

	var cctv model.Cctv
	if err := result.Decode(&cctv); err != nil {
		return nil
	}

	return &cctv
}

func (r *Repository) Update(ctx context.Context, cctv *model.Cctv) error {
	_, err := r.db.Collection(COLLECTION_NAME).UpdateOne(ctx, bson.D{{Key: "id", Value: cctv.Id}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: cctv.Name},
			{Key: "status", Value: cctv.Status},
			{Key: "buildingId", Value: cctv.BuildingId},
			{Key: "events", Value: cctv.Events},
		}}})
	if err != nil {
		return err
	}

	return nil
}
