package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/building/services"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	type testCase struct {
		name         string
		reqBody      string
		createError  error
		expectCode   int
		expectErrors bool
	}

	testCases := []testCase{
		{
			name:         "should return error when failed to bind request body",
			reqBody:      `{"name":"Rented House 2","lat":"invalid type","long":"invalid type"}`,
			createError:  nil,
			expectCode:   400,
			expectErrors: false,
		},
		{
			name:         "should return validation error",
			reqBody:      `{"name":"Rented House 2","lat":-6.4401823}`,
			createError:  nil,
			expectCode:   400,
			expectErrors: true,
		},
		{
			name:         "should return status created",
			reqBody:      `{"name":"Rented House 2","lat":-6.4401823,"long":106.7229039}`,
			createError:  nil,
			expectCode:   http.StatusCreated,
			expectErrors: false,
		},
		{
			name:         "should return error from service",
			reqBody:      `{"name":"Rented House 2","lat":-6.4401823,"long":106.7229039}`,
			createError:  service_error.NewConflictError("some error"),
			expectCode:   http.StatusConflict,
			expectErrors: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := echo.New()

			service := services.NewMockIService(t)
			service.EXPECT().Create(mock.Anything, mock.Anything).
				Return(dto.CreateResponse{}, testCase.createError).Maybe()

			handler := &Handler{
				service: service,
			}
			e.POST("/buildings", handler.Create)

			req := httptest.NewRequest(http.MethodPost, "/buildings", strings.NewReader(testCase.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.Create(c)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectCode, rec.Code)

			var body map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &body)

			if testCase.expectErrors {
				assert.NotNil(t, body["errors"])
			} else {
				assert.Nil(t, body["errors"])
			}
		})
	}
}
