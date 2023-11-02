package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/dto"
    "github.com/go-playground/validator/v10"
  "Paginations/dao"
  )

// @Summary      PUT Customer input: Customer
// @Description  PUT Customer API
// @Tags         PUT Customer - Customer
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Customer false "string collection" 
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateCustomer [PUT]

    func UpdateCustomerApi(c *fiber.Ctx) error {



    inputObj := dto.Customer{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateCustomer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}