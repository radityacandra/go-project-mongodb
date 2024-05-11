package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/repository"
)

type IService interface {
	Create(ctx context.Context, req dto.CreateRequest) (dto.CreateResponse, error)
	List(ctx context.Context, req dto.ListRequest) (dto.MetaResponse, []dto.ListResponse)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}
