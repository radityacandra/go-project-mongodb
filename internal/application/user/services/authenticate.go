package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/internal/application/user/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/radityacandra/go-project-mongodb/pkg/password"
)

func (s *Service) Authenticate(ctx context.Context, req dto.AuthenticateRequest) (dto.AuthenticateResponse, error) {
	user := s.repository.GetByUsername(ctx, req.Username)
	if user == nil {
		return dto.AuthenticateResponse{}, service_error.NewGeneralError("invalid username or password")
	}

	if !password.ValidatePassword(req.Password, user.Password) {
		return dto.AuthenticateResponse{}, service_error.NewGeneralError("invalid username or password")
	}

	// generate jwt
	jwt, exp, err := jwt.BuildToken(map[string]interface{}{
		"userId":   user.Id,
		"username": user.Username,
	})
	if err != nil {
		return dto.AuthenticateResponse{}, service_error.NewGeneralError("failed to build access token")
	}

	return dto.AuthenticateResponse{
		AccessToken: jwt,
		ExpiredAt:   exp,
	}, nil
}
