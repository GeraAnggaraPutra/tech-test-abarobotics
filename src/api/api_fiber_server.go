package api

import (
	"context"
	"fmt"
	"time"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"abarobotics-test/src/kernel"
	"abarobotics-test/toolkit/config"
	"abarobotics-test/toolkit/logger"
)

func RunFiberServer(ctx context.Context, k *kernel.Kernel) {
	cfg := config.NewRuntimeConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler:          errorHandler(),
		DisableStartupMessage: true,
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	if cfg.Prometheus {
		prometheus := fiberprometheus.New(cfg.Name)
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Register routes
	routes(app, k)

	// graceful shutdown
	go func() {
		<-ctx.Done()

		<-time.After(cfg.ShutdownWaitDuration)

		if err := app.Shutdown(); err != nil {
			logger.PrintError(err, "ERROR shutdown server")
		}
	}()

	logger.PrintInfo("serving REST HTTP server", "config", cfg)

	if err := app.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil {
		logger.PrintError(err, "starting http server")
	}
}
