package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	Create(ctx context.Context, cctv *model.Cctv) error
	FindByUserIdAndName(ctx context.Context, userId, name string) *model.Cctv
	FindById(ctx context.Context, userId, id string) *model.Cctv
	Update(ctx context.Context, cctv *model.Cctv) error
}

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
