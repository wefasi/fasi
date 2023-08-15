package handler

import (
	"errors"
	"log"
	"mime"
	"path"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wefasi/fasi/cmd/fasi/infraestructure"
)

func getSiteRelease(site string) (string, error) {
	storage := infraestructure.GetS3()
	cache := infraestructure.GetCache()
	releaseKey := filepath.Join(site, "release")
	release, err := cache.Get(releaseKey)

	log.Println("release: ", release)
	if err != nil {
		log.Println("[DEBUG] release miss ", err)
		release, err = storage.Get(releaseKey)

		if err != nil {
			log.Println("relase file err: ", err)
			return "", errors.New("not found")
		}

		cache.Put(releaseKey, release)
	}

	return release, nil
}

func Proxy(c *fiber.Ctx) error {
	site := c.Hostname()

	customHost := c.Locals("host")
	if customHost != nil {
		site = customHost.(string)
	}

	release, _ := getSiteRelease(site)
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
	cache := infraestructure.GetCache()

	content, err := cache.Get(resource)
	cacheHit := "HIT"

	if err != nil {
		content, err = storage.Get(resource)
		cacheHit = "MISS"
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
		} else {
			cache.Put(resource, content)
		}
	}

	contentType := mime.TypeByExtension(ext)
	c.Set(fiber.HeaderContentType, contentType)
	c.Set(fiber.HeaderCacheControl, "public, max-age=60")

	log.Println("[DEBUG]", cacheHit, "-", resource)

	return c.Status(status).SendString(content)
}
