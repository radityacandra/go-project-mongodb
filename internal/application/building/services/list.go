package services

import (
	"context"
	"math"

	"github.com/radityacandra/go-project-mongodb/internal/application/building/dto"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func (s *Service) List(ctx context.Context, req dto.ListRequest) (dto.MetaResponse, []dto.ListResponse) {
	tokenData := ctx.Value(jwt.TOKEN_DATA).(map[string]string)
	userId := tokenData["userId"]

	buildings, total := s.repository.FindBuildingByUserId(ctx, userId, req.Page, req.PerPage, req.Sort, req.Order)

	var result []dto.ListResponse
	for _, building := range buildings {
		result = append(result, dto.ListResponse{
			BuildingId: building.Id,
			Name:       building.Name,
			Latitude:   building.Latitude,
			Longitude:  building.Longitude,
		})
	}

	return dto.MetaResponse{
		MaxPage: int(math.Ceil(float64(total) / float64(req.PerPage))),
		PerPage: req.PerPage,
		Page:    req.Page,
	}, result
}
