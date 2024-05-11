package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) Update(ctx context.Context, req dto.UpdateRequest) (dto.UpdateResponse, error) {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]

	cctv := s.repository.FindById(ctx, userId, req.Id)
	if cctv == nil {
		return dto.UpdateResponse{}, service_error.NewNotFoundError("cctv not found")
	}

	cctv.UpdateName(req.Name)
	cctv.UpdateStatus(req.Status)
	cctv.UpdateBuildingId(req.BuildingId)
	if err := s.repository.Update(ctx, cctv); err != nil {
		return dto.UpdateResponse{}, service_error.NewDatabaseError("failed to save data")
	}

	return dto.UpdateResponse{
		Id:         cctv.Id,
		Name:       cctv.Name,
		Status:     cctv.Status,
		BuildingId: cctv.BuildingId,
	}, nil
}
