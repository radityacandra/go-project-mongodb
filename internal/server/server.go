package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	"go.uber.org/zap"
)

func InitServer(ctx context.Context, deps *types.Dependency) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())

	deps.Echo = e

	InitRoutes(deps)

	deps.Logger.Info("Web server ready", zap.Int("port", 8080))

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			deps.Logger.Error("Failed to start web server", zap.Error(err))
		}
	}()

	<-ctx.Done()
	deps.Logger.Info("Gracefully shutting down web server...")
	e.Shutdown(ctx)
	deps.Logger.Info("web server shutted down")
}
