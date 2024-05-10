package server

import (
	userHandler "github.com/radityacandra/go-project-mongodb/internal/application/user/handler"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
)

func InitRoutes(deps *types.Dependency) {
	userHandler := userHandler.NewHandler(deps)
	deps.Echo.POST("/users", userHandler.Register)
	deps.Echo.POST("/users/auth", userHandler.Authenticate)
}
