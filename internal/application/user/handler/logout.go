package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/radityacandra/go-project-mongodb/pkg/wrapper/error_wrapper"
)

func (h *Handler) Logout(ctx echo.Context) error {
	err := h.service.Logout(context.WithValue(ctx.Request().Context(), jwt.TOKEN_DATA, ctx.Get(jwt.CONTEXT_KEY)))
	if err == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return error_wrapper.GlobalErrorHandler(err, ctx)
}
