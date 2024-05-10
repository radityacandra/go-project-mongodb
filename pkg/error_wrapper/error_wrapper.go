package error_wrapper

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
)

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Timestamp int64         `json:"timestamp"`
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	Errors    []ErrorDetail `json:"errors"`
}

func GlobalErrorHandler(err error, c echo.Context) error {
	switch typedErr := err.(type) {
	case *echo.HTTPError:
		return c.JSON(typedErr.Code, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      typedErr.Code,
			Message:   err.Error(),
		})
	case validator.ValidationErrors:
		errorResponse := ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      111,
			Message:   err.Error(),
		}

		for _, err := range typedErr {
			errorResponse.Errors = append(errorResponse.Errors, ErrorDetail{
				Field:   err.Field(),
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	case *service_error.DatabaseError:
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      500,
			Message:   err.Error(),
		})
	case *service_error.ConflictError:
		return c.JSON(http.StatusConflict, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      409,
			Message:   err.Error(),
		})
	case *service_error.GeneralError:
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      400,
			Message:   err.Error(),
		})
	case *service_error.NotFoundError:
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      404,
			Message:   err.Error(),
		})
	default:
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Timestamp: time.Now().Unix(),
			Code:      500,
			Message:   typedErr.Error(),
		})
	}
}
