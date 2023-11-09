package apiHandlers

import (
	"Paginations/api"

	"encoding/gob"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	store := session.New()
	//encryptcookie.New(encryptcookie.Config{
	//	Key: "secret-thirty-2-character-string",
	//})
	gob.Register(map[string]interface{}{})

	api := app.Group("/Paginations/api")
	check := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	DefaultMappings(check, store)
	RouteMappings(api, store)
}

func RouteMappings(cg fiber.Router, store *session.Store) {
	cg.Post("/CreateCustomer", api.CreateCustomerApi)
	cg.Put("/UpdateCustomer", api.UpdateCustomerApi)
	cg.Delete("/DeleteCustomer", api.DeleteCustomerApi)
	cg.Get("/FindCustomer", api.FindCustomerApi)
	cg.Get("/FindallCustomer", api.FindallCustomerApi)
	cg.Post("/CreateDriver", api.CreateDriverApi)
	cg.Put("/UpdateDriver", api.UpdateDriverApi)
	cg.Delete("/DeleteDriver", api.DeleteDriverApi)
	cg.Get("/FindDriver", api.FindDriverApi)
	cg.Get("/FindallDriver", api.FindallDriverApi)
	cg.Get("/findFileCount", api.FindfFileCount)

}

func DefaultMappings(cg fiber.Router, store *session.Store) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "service is up and running", "version": "1.0"})
	})
}
