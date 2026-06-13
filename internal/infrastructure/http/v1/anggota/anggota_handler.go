package anggota

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AnggotaHandler struct {
	useCase ports.IAnggotaService
}

func NewAnggotaHandler(uc ports.IAnggotaService) *AnggotaHandler {
	return &AnggotaHandler{uc}
}

func (h *AnggotaHandler) GetAllAnggota(c *fiber.Ctx) error {
	data, err := h.useCase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Gagal mengambil data anggota: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get all anggota",
		"data":  data,
	})
}
