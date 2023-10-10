package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func Register(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			reqid := uuid.New().String()
			log.Logger = log.Logger.With().Interface("request", reqid).Logger()
			return reqid
		},
	}))
	app.Use(logging())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowHeaders: "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept, Origin, Cache-Control, X-Requested-With, Content-Disposition, x-org-id",
		ExposeHeaders: "Content-Disposition, X-Cache",
	}))
}
