package api

import (
	"github.com/gofiber/fiber/v2"
	"sellboot/users"
)

func setUpUserRoutes(web *users.Web, app *fiber.App) {
	router := app.Group("/users")

	router.Post("/", web.RegistrationAdminHandler)
}
