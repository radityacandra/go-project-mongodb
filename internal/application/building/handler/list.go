package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/radityacandra/go-project-mongodb/pkg/wrapper/error_wrapper"
	"github.com/radityacandra/go-project-mongodb/pkg/wrapper/pagination_wrapper"
)

func (h *Handler) List(ctx echo.Context) error {
	var req dto.ListRequest = dto.ListRequest{
		Page:    1,
		PerPage: 10,
		Sort:    "name",
		Order:   "asc",
	}

	if err := ctx.Bind(&req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return error_wrapper.GlobalErrorHandler(err, ctx)
	}

	meta, data := h.service.List(context.WithValue(ctx.Request().Context(), jwt.TOKEN_DATA, ctx.Get(jwt.CONTEXT_KEY)), req)
	return ctx.JSON(http.StatusOK, pagination_wrapper.PaginationWrapper(meta.Page, meta.PerPage, meta.MaxPage, data))
}
