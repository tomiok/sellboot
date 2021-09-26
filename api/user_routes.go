package api

import (
	"github.com/gofiber/fiber/v2"
	"sellboot/users"
)

func setUpUserRoutes(web *users.Web, app *fiber.App) {
	router := app.Group("/users")

	router.Post("/admin", web.RegistrationAdminHandler)
	router.Post("/investor", web.RegistrationInvestorHandler)
	router.Post("/company", web.RegistrationCompanyHandler)

	router.Post("/login", web.AuthorizationHandler)
	router.Get("/profile", web.UserProfileHandler)
}
