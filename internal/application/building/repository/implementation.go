package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	"go.mongodb.org/mongo-driver/bson"
)

const COLLECTION_NAME = "buildings"

func (r *Repository) Create(ctx context.Context, building *model.Building) error {
	_, err := r.db.Collection(COLLECTION_NAME).InsertOne(ctx, building)
	return err
}

func (r *Repository) FindBuildingByUserIdAndName(ctx context.Context, userId, name string) *model.Building {
	var building model.Building
	result := r.db.Collection(COLLECTION_NAME).
		FindOne(ctx, bson.D{{Key: "userId", Value: userId}, {Key: "name", Value: name}})

	if result.Err() != nil {
		return nil
	}

	if err := result.Decode(&building); err != nil {
		return nil
	}

	return &building
}

func (r *Repository) FindBuildingByUserId(ctx context.Context, userId string, page, pageSize int, sortBy, order string) ([]model.Building, int) {
	var buildings []model.Building

	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	}

	filter := bson.D{{Key: "userId", Value: userId}}
	result, err := r.db.Collection(COLLECTION_NAME).
		Aggregate(ctx, bson.A{
			bson.D{{Key: "$match", Value: filter}},
			bson.D{{Key: "$sort", Value: bson.D{{Key: sortBy, Value: sortOrder}}}},
			bson.D{{Key: "$skip", Value: pageSize * (page - 1)}},
			bson.D{{Key: "$limit", Value: pageSize}},
		})
	if err != nil {
		return buildings, 0
	}

	result.All(ctx, &buildings)

	resultTotal, err := r.db.Collection(COLLECTION_NAME).Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: filter}},
		bson.D{{Key: "$count", Value: "total"}},
	})
	if err != nil {
		return buildings, 0
	}

	var totalDto []dto.CountDto
	resultTotal.All(ctx, &totalDto)

	var total int
	if len(totalDto) > 0 {
		total = totalDto[0].Total
	}

	return buildings, total
}
