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
	_ "github.com/mhrndiva/ws-ats-714220050/docs"

	"github.com/gofiber/fiber/v2"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/mhrndiva
// @contact.email 714220050@std.ulbi.ac.id

// @host ws-ats-714220050-cc57ecdf5b73.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
