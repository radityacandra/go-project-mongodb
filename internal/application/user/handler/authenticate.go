package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/user/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/error_wrapper"
)

func (h *Handler) Authenticate(ctx echo.Context) error {
	var req dto.AuthenticateRequest

	if err := ctx.Bind(&req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	res, err := h.service.Authenticate(ctx.Request().Context(), req)
	if err == nil {
		return ctx.JSON(http.StatusOK, res)
	}

	return error_wrapper.GlobalErrorHandler(err, ctx)
}
