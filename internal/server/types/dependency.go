package types

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Dependency struct {
	Logger *zap.Logger
	DB     *mongo.Database
	Echo   *echo.Echo
}

func NewDependency(logger *zap.Logger, db *mongo.Database) *Dependency {
	return &Dependency{
		Logger: logger,
		DB:     db,
	}
}
