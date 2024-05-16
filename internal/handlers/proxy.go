package handlers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/jkdv-systeme/kyasshu/internal/config"
	"github.com/jkdv-systeme/kyasshu/internal/responses"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"net/url"
	"strings"
)

func proxy(c *fiber.Ctx) error {
	logger := c.Locals("logger").(zerolog.Logger)

	path := c.Params("+")

	if path == "" {
		return responses.NewError(fiber.StatusBadRequest, "path must be specified")
	}

	path, err := url.QueryUnescape(strings.TrimPrefix(path, "/"))

	if err != nil {
		logger.Error().Err(err).Msg("failed to unescape path")
		return responses.NewError(fiber.StatusBadRequest, "path must be specified")
	}

	logger.Info().Str("host", c.Hostname()).Msgf("proxying request for %s", path)

	var domains config.DomainConfig

	err = viper.UnmarshalKey("domains", &domains)

	if err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal domains")
		return responses.NewError(fiber.StatusInternalServerError, "could not download file")
	}

	tenantConfig, ok := domains[c.Hostname()]

	if !ok {
		return responses.NewError(fiber.StatusBadRequest, "domain not registered")
	}

	cfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials(tenantConfig.AccessKey, tenantConfig.SecretKey, ""),
		Endpoint:    aws.String(tenantConfig.Endpoint),
		Region:      aws.String(tenantConfig.Region),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create aws session")
		return responses.NewError(fiber.StatusInternalServerError, "could not download file")
	}

	client := s3.New(sess)

	out, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(tenantConfig.Bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed to get object")
		return responses.NewError(fiber.StatusNotFound, "file not found")
	}

	//log.Info().Interface("out", out).Msgf("out: %v", out)

	c.Set("Content-Type", *out.ContentType)
	//c.Set("Content-Disposition", "inline; filename="+filepath.Base(path))

	c.Set("Cross-Origin-Resource-Policy", "cross-origin")
	c.Set("Cross-Origin-Opener-Policy", "cross-origin")

	return c.SendStream(out.Body, int(*out.ContentLength))
}
