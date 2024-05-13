package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	Create(ctx context.Context, building *model.Building) error
	FindBuildingByUserId(ctx context.Context, userId string,
		page, pageSize int, sortBy, order string, keyword *string) ([]model.Building, int)
	FindBuildingByUserIdAndName(ctx context.Context, userId, name string) *model.Building
}

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
