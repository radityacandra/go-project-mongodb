package jwt

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/user/repository"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	jwtTypes "github.com/radityacandra/go-project-mongodb/pkg/jwt/types"
	"github.com/radityacandra/go-project-mongodb/pkg/wrapper/error_wrapper"
)

type tokenData string

type JwtMiddleware struct {
	repository repository.IRepository
}

func NewJwtMiddlewre(deps *types.Dependency) *JwtMiddleware {
	repository := repository.NewRepository(deps.DB)

	return &JwtMiddleware{
		repository: repository,
	}
}

func (m *JwtMiddleware) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := AuthorizeToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return error_wrapper.GlobalErrorHandler(err, c)
		}

		if data["userId"] == "" || data["sessionId"] == "" {
			return error_wrapper.GlobalErrorHandler(jwtTypes.NewAuthorizationError("invalid token"), c)
		}

		user := m.repository.GetUserBySessionId(c.Request().Context(), data["userId"], data["sessionId"])
		if user == nil {
			return error_wrapper.GlobalErrorHandler(jwtTypes.NewAuthorizationError("no active session is found"), c)
		}

		c.Set(CONTEXT_KEY, data)

		return next(c)
	}
}
