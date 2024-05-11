package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) Create(ctx context.Context, req dto.CreateRequest) (dto.CreateResponse, error) {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]

	if existingBuilding := s.repository.FindBuildingByUserIdAndName(ctx, userId, req.Name); existingBuilding != nil {
		return dto.CreateResponse{}, service_error.NewConflictError("existing building with the same name already exist")
	}

	building := model.NewBuilding(userId, req.Name, req.Latitude, req.Longitude)
	if err := s.repository.Create(ctx, building); err != nil {
		return dto.CreateResponse{}, service_error.NewDatabaseError("failed to save building")
	}

	return dto.CreateResponse{
		BuildingId: building.Id,
		Name:       building.Name,
		Latitude:   building.Latitude,
		Longitude:  building.Longitude,
	}, nil
}
