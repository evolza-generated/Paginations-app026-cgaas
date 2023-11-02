package api

import (
  "Paginations/utils"
  "github.com/gofiber/fiber/v2"

  "Paginations/dto"
    "github.com/go-playground/validator/v10"
  "Paginations/dao"
  )

// @Summary      PUT Driver input: Driver
// @Description  PUT Driver API
// @Tags         PUT Driver - Driver
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Driver false "string collection" 
// @Success      200  {array}   dto.Model_Driver "Status OK"
// @Success      202  {array}   dto.Model_Driver "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateDriver [PUT]

    func UpdateDriverApi(c *fiber.Ctx) error {



    inputObj := dto.Driver{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateDriver(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}