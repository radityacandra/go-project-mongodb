package services

import (
	"context"
	"errors"
	"testing"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/building/repository"
	"github.com/radityacandra/go-project-mongodb/pkg/errors/service_error"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	type testCase struct {
		name                        string
		req                         dto.CreateRequest
		findBuildingByUserIdAndName *model.Building
		create                      error
		expectDto                   dto.CreateResponse
		expectErr                   error
	}

	testCases := []testCase{
		{
			name: "should return error if same building is found on the same user",
			req: dto.CreateRequest{
				Name: "test",
			},
			findBuildingByUserIdAndName: &model.Building{},
			expectDto:                   dto.CreateResponse{},
			expectErr:                   service_error.NewConflictError("existing building with the same name already exist"),
		},
		{
			name: "should return error if failed to save to db",
			req: dto.CreateRequest{
				Name: "test",
			},
			findBuildingByUserIdAndName: nil,
			create:                      errors.New("some error"),
			expectDto:                   dto.CreateResponse{},
			expectErr:                   service_error.NewDatabaseError("failed to save building"),
		},
		{
			name: "should return success",
			req: dto.CreateRequest{
				Name:      "test",
				Latitude:  1,
				Longitude: 2,
			},
			findBuildingByUserIdAndName: nil,
			create:                      nil,
			expectDto: dto.CreateResponse{
				Name:      "test",
				Latitude:  1,
				Longitude: 2,
			},
			expectErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), jwt.TOKEN_DATA, map[string]string{
				"userId": "userId1",
			})

			repository := repository.NewMockIRepository(t)
			repository.EXPECT().
				FindBuildingByUserIdAndName(ctx, "userId1", testCase.req.Name).
				Return(testCase.findBuildingByUserIdAndName)

			repository.EXPECT().
				Create(ctx, mock.Anything).
				Return(testCase.create).Maybe()

			service := NewService(repository)
			res, err := service.Create(ctx, testCase.req)

			if res.Name != "" {
				assert.NotEqual(t, "", res.BuildingId)
				res.BuildingId = ""
			}
			assert.Equal(t, testCase.expectDto, res)
			assert.Equal(t, testCase.expectErr, err)
		})
	}
}
