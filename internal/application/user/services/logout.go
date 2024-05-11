package services

import (
	"context"

	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) Logout(ctx context.Context) error {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]
	sessionId := tokenData["sessionId"]

	return s.repository.RevokeSession(ctx, userId, sessionId)
}
