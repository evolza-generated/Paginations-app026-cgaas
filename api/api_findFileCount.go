package api

import (
	"Paginations/dao"
	"Paginations/utils"

	"github.com/gofiber/fiber/v2"
)

func FindfFileCount(c *fiber.Ctx) error {
	count, err := dao.FindFileCountDao()
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(fiber.Map{
		"count": count,
	})
}
