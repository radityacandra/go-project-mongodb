package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/user/model"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/password"
)

func (s *Service) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// check conflicting username
	if existingUser := s.repository.GetByUsername(ctx, req.Username); existingUser != nil {
		return dto.RegisterResponse{}, service_error.NewConflictError("existing user with the same username already exist")
	}

	var hashed string
	if password, err := password.HashPassword(req.Password); err != nil {
		return dto.RegisterResponse{}, service_error.NewGeneralError("failed to create user")
	} else {
		hashed = password
	}

	user := model.NewUser(req.Username, req.PhoneNumber, hashed)

	if err := s.repository.Save(ctx, user); err != nil {
		return dto.RegisterResponse{}, service_error.NewDatabaseError("failed to save user")
	}

	return dto.RegisterResponse{
		UserId:      user.Id,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
