package services

import (
	"context"

	"github.com/google/uuid"
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
	sessionId := uuid.NewString()
	jwt, exp, err := jwt.BuildToken(map[string]interface{}{
		"userId":    user.Id,
		"sessionId": sessionId,
		"username":  user.Username,
	})
	if err != nil {
		return dto.AuthenticateResponse{}, service_error.NewGeneralError("failed to build access token")
	}

	user.AddSession(sessionId, exp)
	if err := s.repository.SaveSession(ctx, user); err != nil {
		return dto.AuthenticateResponse{}, service_error.NewDatabaseError("failed to save session")
	}

	return dto.AuthenticateResponse{
		AccessToken: jwt,
		ExpiredAt:   exp,
	}, nil
}
