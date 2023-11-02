package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/dao"
  )

// @Summary      GET Driver input: Driver
// @Description  GET Driver API
// @Tags         GET Driver - Driver
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Driver "Status OK"
// @Success      202  {array}   dto.Model_Driver "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindDriver [GET]

    func FindDriverApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindDriverbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}