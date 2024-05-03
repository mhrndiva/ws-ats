package controller

import (
	"github.com/mhrndiva/kemahasiswaan"
	"github.com/gofiber/fiber/v2"

)

func Homepage(c *fiber.Ctx) error {
	ipaddr := _714220050.GetAllMahasiswa()
	return c.JSON(ipaddr)
}

