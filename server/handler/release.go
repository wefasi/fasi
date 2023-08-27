package handler

import (
	"log"
	"os"
	"path"

	"github.com/wefasi/fasi/server/infraestructure"

	"github.com/gofiber/fiber/v2"
)

func SetSiteRelease(c *fiber.Ctx) error {
	apiKey := c.GetReqHeaders()["Authorization"]
	if apiKey != os.Getenv("API_KEY") {
		return c.Status(403).SendString("Forbidden. Invalid API Key")
	}

	site := c.Params("site")
	if site == "" {
		return c.Status(400).SendString("missing site")
	}

	release := c.Params("release")
	if site == "" {
		return c.Status(400).SendString("missing release")
	}

	cache := infraestructure.GetCache()
	cache.Put(path.Join(site, "release"), release)

	log.Println("[DEBUG] set-site-release:", site, release)

	return c.SendStatus(200)
}
