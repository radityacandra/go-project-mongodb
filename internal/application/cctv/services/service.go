package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/repository"
)

type IService interface {
	Create(ctx context.Context, req dto.CreateRequest) (dto.CreateResponse, error)
	Update(ctx context.Context, req dto.UpdateRequest) (dto.UpdateResponse, error)
	ReportEvent(ctx context.Context, req dto.ReportEventRequest) (dto.ReportEventResponse, error)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}
