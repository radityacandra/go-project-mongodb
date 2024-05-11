package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/model"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) Create(ctx context.Context, req dto.CreateRequest) (dto.CreateResponse, error) {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]

	// TODO: cctv name check

	cctv := model.NewCctv(userId, req.Name, req.DeviceId)
	if err := s.repository.Create(ctx, cctv); err != nil {
		return dto.CreateResponse{}, service_error.NewDatabaseError("failed to save cctv")
	}

	return dto.CreateResponse{
		Id:         cctv.Id,
		Name:       cctv.Name,
		DeviceId:   cctv.DeviceId,
		Status:     cctv.Status,
		LastUpdate: cctv.LastUpdate,
	}, nil
}
