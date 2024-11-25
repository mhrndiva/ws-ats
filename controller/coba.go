package controller

import (
	"fmt"
	//"github.com/mhrndiva/kemahasiswaan"
	"github.com/gofiber/fiber/v2"
	inimodel "github.com/mhrndiva/kemahasiswaan/model"
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
// // GetPresensi godoc
// // @Summary Get All Data Presensi.
// // @Description Mengambil semua data presensi.
// // @Tags Presensi
// // @Accept json
// // @Produce json
// // @Success 200 {object} Presensi
// // @Router /presensi [get]
// func GetPresensi(c *fiber.Ctx) error {
// 	ps := cek.GetAllPresensi()
// 	return c.JSON(ps)
// }

// func GetMatkul (c *fiber.Ctx) error {
// 	ps := cek.GetAllMatkul()
// 	return c.JSON(ps)
// }
// // InsertDataPresensi godoc
// // @Summary Insert data presensi.
// // @Description Input data presensi.
// // @Tags Presensi
// // @Accept json
// // @Produce json
// // @Param request body ReqPresensi true "Payload Body [RAW]"
// // @Success 200 {object} Presensi
// // @Failure 400
// // @Failure 500
// // @Router /devi [post]
// func InsertDataPresensi(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var presensi inimodel.Presensi
// 	if err := c.BodyParser(&presensi); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	insertedID, err := cek.InsertPresensi(db, "presensi",
// 		presensi.Npm,
// 		presensi.Matkul,
// 		presensi.Biodata,
// 		presensi.Checkin)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }
// // GetPresensiID godoc
// // @Summary Get By ID Data Presensi.
// // @Description Ambil per ID data presensi.
// // @Tags Presensi
// // @Accept json
// // @Produce json
// // @Param id path string true "Masukan ID"
// // @Success 200 {object} Presensi
// // @Failure 400
// // @Failure 404
// // @Failure 500
// // @Router /presensi/{id} [get]
// func GetPresensiID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Wrong parameter",
// 		})
// 	}
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "Invalid id parameter",
// 		})
// 	}
// 	ps, err := cek.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 				"status":  http.StatusNotFound,
// 				"message": fmt.Sprintf("No data found for id %s", id),
// 			})
// 		}
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": fmt.Sprintf("Error retrieving data for id %s", id),
// 		})
// 	}
// 	return c.JSON(ps)
// }

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqPresensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

// 	// Parse the request body into a Presensi object
// 	var presensi inimodel.Presensi
// 	if err := c.BodyParser(&presensi); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}

// 	// Call the UpdatePresensi function with the parsed ID and the Presensi object
// 	err = cek.UpdatePresensi(db, "presensi",
// 		objectID,
// 		presensi.Npm,
// 		presensi.Matkul,
// 		presensi.Biodata,
// 		presensi.Checkin)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":  http.StatusOK,
// 		"message": "Data successfully updated",
// 	})
// }

// // DeletePresensiByID godoc
// // @Summary Delete data presensi.
// // @Description Hapus data presensi.
// // @Tags Presensi
// // @Accept json
// // @Produce json
// // @Param id path string true "Masukan ID"
// // @Success 200
// // @Failure 400
// // @Failure 500
// // @Router /delete/{id} [delete]
// func DeletePresensiByID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Wrong parameter",
// 		})
// 	}

// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "Invalid id parameter",
// 		})
// 	}

// 	err = cek.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": fmt.Sprintf("Error deleting data for id %s", id),
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":  http.StatusOK,
// 		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
// 	})
// }