package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/functions"
    "Paginations/dto"
    "github.com/go-playground/validator/v10"
  "Paginations/dao"
  )

// @Summary      POST Customer input: Customer
// @Description  POST Customer API
// @Tags         POST Customer - Customer
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Customer false "string collection" 
// @Success      200  {array}   dto.Model_Customer "Status OK"
// @Success      202  {array}   dto.Model_Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateCustomer [POST]

    func CreateCustomerApi(c *fiber.Ctx) error {



    inputObj := dto.Customer{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
if err := functions.UniqueCheck(inputObj, "Customers", []string{ "customerID", "email", "phone"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateCustomer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}