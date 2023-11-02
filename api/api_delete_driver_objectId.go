package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/dao"
  )

// @Summary      DELETE Driver input: Driver
// @Description  DELETE Driver API
// @Tags         DELETE Driver - Driver
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Driver "Status OK"
// @Success      202  {array}   dto.Model_Driver "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteDriver [DELETE]

    func DeleteDriverApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteDriver(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}