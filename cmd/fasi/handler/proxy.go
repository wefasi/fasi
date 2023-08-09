package handler

import (
	"log"
	"mime"
	"path"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wefasi/fasi/cmd/fasi/infraestructure"
)

func Proxy(c *fiber.Ctx) error {
	site := c.Hostname()
	release := ""
	customHost := c.Locals("host")
	if customHost != nil {
		site = customHost.(string)
	}

	customRelease := c.Locals("release")
	if customRelease != nil {
		release = customRelease.(string)
	}

	request_path := c.Path()
	if strings.HasSuffix(request_path, "/") {
		request_path += "index"
	}

	if path.Ext(request_path) == "" {
		request_path = request_path + "." + "html"
	}

	ext := path.Ext(request_path)

	status := 200
	resource := filepath.Join(site, release, request_path)
	storage := infraestructure.GetS3()
	content, err := storage.Get(resource)
	if err != nil {
		if err.Error() == "not found" {
			status = 404
			content, err = storage.Get(path.Join(site, release, "404.html"))
		}

		if err != nil {
			log.Println("[ERROR] " + err.Error())
			status = 500
			content = "oops, something went wrong :("
		}
	}

	contentType := mime.TypeByExtension(ext)
	c.Set(fiber.HeaderContentType, contentType)

	// TODO improve checking extension
	c.Set(fiber.HeaderCacheControl, "public, max-age=60")

	return c.Status(status).SendString(content)
}
