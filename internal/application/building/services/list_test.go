package services

import (
	"context"
	"testing"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/building/repository"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ctx := context.WithValue(context.Background(), jwt.TOKEN_DATA, map[string]string{
		"userId": "userId1",
	})

	keyword := "test"

	req := dto.ListRequest{
		Page:    2,
		PerPage: 200,
		Sort:    "name",
		Order:   "asc",
		Keyword: &keyword,
	}

	repository := repository.NewMockIRepository(t)
	repository.EXPECT().
		FindBuildingByUserId(ctx, "userId1", req.Page, req.PerPage, req.Sort, req.Order, req.Keyword).
		Return([]model.Building{{
			Id: "1",
		}}, 100)

	service := NewService(repository)
	meta, res := service.List(ctx, req)

	assert.Equal(t, 1, meta.MaxPage)
	assert.Equal(t, 2, meta.Page)
	assert.Equal(t, 200, meta.PerPage)

	assert.Equal(t, 1, len(res))
}
