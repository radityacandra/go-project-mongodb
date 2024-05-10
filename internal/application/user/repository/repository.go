package repository

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	Save(ctx context.Context, user *model.User) error
	GetByUsername(ctx context.Context, userName string) *model.User
}

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}
