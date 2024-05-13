package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/radityacandra/go-project-mongodb/pkg/wrapper/error_wrapper"
)

func (h *Handler) Create(ctx echo.Context) error {
	var req dto.CreateRequest

	if err := ctx.Bind(&req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	res, err := h.service.Create(context.WithValue(ctx.Request().Context(), jwt.TOKEN_DATA, ctx.Get(jwt.CONTEXT_KEY)), req)
	if err == nil {
		return ctx.JSON(http.StatusCreated, res)
	}

	return error_wrapper.GlobalErrorHandler(err, ctx)
}
