package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkdv-systeme/kyasshu/internal/responses"
	"github.com/rs/zerolog/log"
	"time"
)

func logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		c.Locals("logger", log.Logger.With().Interface("request", c.Locals("requestid")).Logger())

		log.Info().Str("method", c.Method()).Str("path", c.Path()).Msg("Received request")
		// Go to next middleware:
		next := c.Next()
		duration := time.Now().Sub(start)
		if next != nil {
			if err, ok := next.(*responses.ServerError); ok {
				log.Info().Interface("request", c.Locals("requestid")).Int("status", err.Status).Str("method", c.Method()).Str("path", c.Path()).Str("duration", duration.String()).Msg("Handled request")
				return err
			}
		} else {
			log.Info().Interface("request", c.Locals("requestid")).Int("status", c.Response().StatusCode()).Str("method", c.Method()).Str("path", c.Path()).Str("duration", duration.String()).Msg("Handled request")
		}
		return next
	}
}
