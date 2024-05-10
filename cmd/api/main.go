package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/radityacandra/go-project-mongodb/internal/server"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	"github.com/radityacandra/go-project-mongodb/pkg/database"
	"github.com/radityacandra/go-project-mongodb/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.LoadLogger()

	logger.Info("starting application, press CTRL + C to gracefully shutdown")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := database.InitDatabase(ctx)
	if err != nil {
		logger.Error("failed to establish db connection", zap.Error(err))
	}

	dependency := types.NewDependency(logger, db.GetDatabase())

	// start web server
	server.InitServer(ctx, dependency)

	<-ctx.Done()
	db.Disconnect(context.Background())
}
