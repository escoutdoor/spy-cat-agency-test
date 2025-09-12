package main

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/app"
	"github.com/escoutdoor/spy-cat-agency-test/internal/config"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/closer"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	if err := config.Load(); err != nil {
		logger.Fatal(ctx, "load config:", err)
	}

	if config.Config().App.IsProd() {
		logger.SetLevel(zap.InfoLevel)
	} else {
		logger.SetLevel(zap.DebugLevel)
	}

	closer.SetShutdownTimeout(config.Config().App.GracefulShutdownTimeout())

	a, err := app.New(ctx)
	if err != nil {
		logger.Fatal(ctx, "init app:", err)
	}

	if err := a.Run(ctx); err != nil {
		logger.Fatal(ctx, "run app:", err)
	}

	closer.Wait()
}
