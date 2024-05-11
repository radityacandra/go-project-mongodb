package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) ReportEvent(ctx context.Context, req dto.ReportEventRequest) (dto.ReportEventResponse, error) {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]

	cctv := s.repository.FindById(ctx, userId, req.Id)
	if cctv == nil {
		return dto.ReportEventResponse{}, service_error.NewNotFoundError("cctv not found")
	}

	event := cctv.AddEvent(req.Name, req.Timestamp, req.Type, req.Source)

	if err := s.repository.Update(ctx, cctv); err != nil {
		return dto.ReportEventResponse{}, service_error.NewDatabaseError("failed to report cctv event")
	}

	return dto.ReportEventResponse{
		EventId:   event.Id,
		Name:      event.Name,
		Timestamp: event.Timestamp,
		Type:      event.Type,
		Source:    event.Source,
	}, nil
}
