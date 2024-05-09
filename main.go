package main

import (
	"log"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/indrariksa/cobapakcage/module"
	"github.com/mhrndiva/ws-ats-714220050/config"
	"github.com/mhrndiva/ws-ats-714220050/url"
	//"github.com/mhrndiva/kemahasiswaan/module"
	//"github.com/mhrndiva/kemahasiswaan/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
