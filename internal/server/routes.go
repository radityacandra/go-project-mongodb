package server

import (
	buildingHandler "github.com/radityacandra/go-project-mongodb/internal/application/building/handler"
	cctvHandler "github.com/radityacandra/go-project-mongodb/internal/application/cctv/handler"
	userHandler "github.com/radityacandra/go-project-mongodb/internal/application/user/handler"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	"github.com/radityacandra/go-project-mongodb/pkg/jwt"
)

func InitRoutes(deps *types.Dependency) {
	jwtMiddleware := jwt.NewJwtMiddlewre(deps)

	userHandler := userHandler.NewHandler(deps)
	deps.Echo.POST("/users", userHandler.Register)
	deps.Echo.POST("/users/auth", userHandler.Authenticate)
	deps.Echo.POST("/users/logout", userHandler.Logout, jwtMiddleware.Authorize)

	buildingHandler := buildingHandler.NewHandler(deps)
	deps.Echo.POST("/buildings", buildingHandler.Create, jwtMiddleware.Authorize)
	deps.Echo.GET("/buildings", buildingHandler.List, jwtMiddleware.Authorize)

	cctvHandler := cctvHandler.NewHandler(deps)
	deps.Echo.POST("/cctvs", cctvHandler.Create, jwtMiddleware.Authorize)
	deps.Echo.PUT("/cctvs/:id", cctvHandler.Update, jwtMiddleware.Authorize)
	deps.Echo.POST("/cctvs/:id/events", cctvHandler.ReportEvent, jwtMiddleware.Authorize)
}
