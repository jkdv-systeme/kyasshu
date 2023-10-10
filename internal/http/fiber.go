package http

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jkdv-systeme/kyasshu/internal/responses"
	"github.com/rs/zerolog/log"
)

func getFiberConfig() fiber.Config {
	return fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			reqId := ctx.Locals("requestid")

			requestID := ""

			if reqId != nil {
				requestID = reqId.(string)
			}

			var serverError *responses.ServerError
			if errors.As(err, &serverError) {
				ctx.Status(serverError.Status)
				err := ctx.JSON(&responses.ErrorResponse{
					Status:    serverError.Status,
					Message:   serverError.Message,
					RequestID: requestID,
				})
				return err
			}
			var validationError *responses.ValidationError
			if errors.As(err, &validationError) {
				ctx.Status(validationError.Status)
				err := ctx.JSON(&responses.ErrorResponse{
					Status:    validationError.Status,
					Message:   validationError.Message,
					RequestID: requestID,
					Fields:    validationError.Fields,
				})
				return err
			}
			log.Error().Err(err).Msg("Unhandled internal server error")
			ctx.Status(fiber.StatusInternalServerError)
			err = ctx.JSON(&responses.ErrorResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Something went wrong",
			})
			return err
		},
	}
}
