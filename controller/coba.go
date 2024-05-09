package controller

import (
	//"github.com/mhrndiva/kemahasiswaan"
	"github.com/gofiber/fiber/v2"
	cek "github.com/mhrndiva/kemahasiswaan/module"
	"github.com/aiteung/musik"
	//"github.com/mhrndiva/ws-ats-714220050/config"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	//"net/http"

)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetMahasiswa(c *fiber.Ctx) error {
	ps := cek.GetAllMahasiswa()
	return c.JSON(ps)
}

