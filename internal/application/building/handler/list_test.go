package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/building/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	type testCase struct {
		name         string
		queryParams  string
		listCalls    int
		listArgs     dto.ListRequest
		expectCode   int
		expectErrors bool
	}

	testCases := []testCase{
		{
			name:        "should return error binding query params",
			queryParams: "?page=asdf",
			listCalls:   0,
			expectCode:  400,
		},
		{
			name:         "should return error query params validation",
			queryParams:  "?page=-1",
			listCalls:    0,
			expectCode:   400,
			expectErrors: true,
		},
		{
			name:      "should return success with default value",
			listCalls: 1,
			listArgs: dto.ListRequest{
				Page:    1,
				PerPage: 10,
				Sort:    "name",
				Order:   "asc",
			},
			expectCode: 200,
		},
		{
			name:        "should return success with default value",
			queryParams: "?page=2&per_page=1&sort=lat&order=desc",
			listCalls:   1,
			listArgs: dto.ListRequest{
				Page:    2,
				PerPage: 1,
				Sort:    "lat",
				Order:   "desc",
			},
			expectCode: 200,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := echo.New()

			service := services.NewMockIService(t)
			service.EXPECT().
				List(mock.Anything, testCase.listArgs).
				Return(dto.MetaResponse{}, []dto.ListResponse{}).
				Maybe()

			handler := &Handler{
				service: service,
			}
			e.GET("/buildings", handler.List)

			req := httptest.NewRequest(http.MethodGet, "/buildings"+testCase.queryParams, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.List(c)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectCode, rec.Code)

			var body map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &body)

			if testCase.expectErrors {
				assert.NotNil(t, body["errors"])
			} else {
				assert.Nil(t, body["errors"])
			}

			if testCase.listCalls == 0 {
				service.AssertNotCalled(t, "List")
			} else {
				service.AssertCalled(t, "List", mock.Anything, testCase.listArgs)
			}
		})
	}
}
