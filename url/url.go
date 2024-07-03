package url

import (
	"github.com/mhrndiva/ws-ats-714220050/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	page.Get("/mahasiswa", controller.GetMahasiswa) //panggil data mahasiswa
	page.Get("/presensi", controller.GetPresensi) //panggil data presensi
	page.Get("/matkul", controller.GetMatkul) //panggil matkul
	page.Post("/insert", controller.InsertDataPresensi)
}
