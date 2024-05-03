package main

import (
	"log"

	"github.com/mhrndiva/WS-ATS-714220050/config"

	"github.com/mhrndiva/kemahasiswaan"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/mhrndiva/ws-maharani2024/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
