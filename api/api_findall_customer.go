package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/dao"
  )

// @Summary      GET Customer input: Customer
// @Description  GET Customer API
// @Tags         GET Customer - Customer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallCustomer [GET]

    func FindallCustomerApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallCustomer()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}