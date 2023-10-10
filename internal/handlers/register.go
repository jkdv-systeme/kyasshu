package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/ssl/check", onDemandTls)
	//app.Get("+", cache.New(cache.Config{
	//	Expiration:   1 * time.Hour,
	//	CacheControl: false,
	//}), proxy)
	app.Get("+", proxy)
}
