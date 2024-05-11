package handler

import (
	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/repository"
	"github.com/radityacandra/go-project-mongodb/internal/application/cctv/services"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
)

type Handler struct {
	service services.IService
}

func NewHandler(deps *types.Dependency) *Handler {
	repository := repository.NewRepository(deps.DB)
	service := services.NewService(repository)

	return &Handler{
		service: service,
	}
}
