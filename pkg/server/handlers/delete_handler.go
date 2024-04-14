package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) DeleteBanner(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))
	if id <= 0 || err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	err = h.DB.DeleteBannerById(uint(id))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusNoContent).JSON(&id)
}