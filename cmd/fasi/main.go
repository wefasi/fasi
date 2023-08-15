package main

import (
	"fmt"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wefasi/fasi/cmd/fasi/handler"
	"github.com/wefasi/fasi/cmd/fasi/infraestructure"
)

func newApp() *fiber.App {
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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	infraestructure.InitCache()
	infraestructure.InitS3()
	app := newApp()

	fmt.Println("🚀 Listening http://localhost:3210")
	err := app.Listen("localhost:3210")
	if err != nil {
		panic(err)
	}
}
