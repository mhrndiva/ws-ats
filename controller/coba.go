package controller

import (
	"fmt"
	//"github.com/mhrndiva/kemahasiswaan"
	"github.com/gofiber/fiber/v2"
	cek "github.com/mhrndiva/kemahasiswaan/module"
	"github.com/aiteung/musik"
	"github.com/mhrndiva/ws-ats-714220050/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"errors"
	"net/http"

)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetMahasiswa(c *fiber.Ctx) error {
	ps := cek.GetAllMahasiswa()
	return c.JSON(ps)
}

func GetMahasiswaID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetMahasiswaFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetAllPresensi()
	return c.JSON(ps)
}

func GetMatkul (c *fiber.Ctx) error {
	ps := cek.GetAllMatkul()
	return c.JSON(ps)
}

