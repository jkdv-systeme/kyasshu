package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jkdv-systeme/kyasshu/internal/config"
	"github.com/jkdv-systeme/kyasshu/internal/responses"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func onDemandTls(c *fiber.Ctx) error {
	domain := c.Query("domain")

	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	var domains config.DomainConfig
	err := viper.UnmarshalKey("domains", &domains)

	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal domains")
		return responses.NewError(fiber.StatusForbidden, "domain not registered")
	}

	_, ok := domains[domain]

	if !ok {
		return responses.NewError(fiber.StatusForbidden, "domain not registered")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
