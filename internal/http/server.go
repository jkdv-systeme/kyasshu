package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jkdv-systeme/kyasshu/internal/handlers"
	"github.com/jkdv-systeme/kyasshu/internal/middleware"
	"github.com/jkdv-systeme/kyasshu/internal/responses"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {
	log.Info().Msg("starting server...")

	// Shutdown signal channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	port := viper.GetInt("server.port")

	// Create fiber app instance
	app := fiber.New(getFiberConfig())

	// Global middleware
	middleware.Register(app)

	handlers.Register(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return responses.NewError(fiber.StatusNotFound, "the requested endpoint does not exist")
	})

	go func() {
		log.Info().Int("port", port).Msg("api is ready and listening")
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Error().Err(err).Msg("error starting api application")
		}
	}()

	<-done

	log.Info().Msg("stopping server...")

	//err := app.Shutdown()
	//if err != nil {
	//	log.Error().Err(err).Msg("failed to shut down api server")
	//}

	log.Info().Msg("server gracefully shut down")

}
