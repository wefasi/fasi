package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wefasi/fasi/server/handler"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(func(c *fiber.Ctx) error {
		if len(os.Args) >= 2 {
			c.Locals("host", os.Args[1])
		}

		if len(os.Args) >= 3 {
			c.Locals("release", os.Args[2])
		}

		return c.Next()
	})

	app.Put("/api/i/site/:site/release/:release", handler.SetSiteRelease)
	app.Get("*", handler.Proxy)

	return app
}
