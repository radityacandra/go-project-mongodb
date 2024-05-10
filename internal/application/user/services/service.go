package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/user/repository"
)

type IService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
	Authenticate(ctx context.Context, req dto.AuthenticateRequest) (dto.AuthenticateResponse, error)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}
