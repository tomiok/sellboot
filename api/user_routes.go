package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"sellboot/users"
)

func setUpUserRoutes(store *session.Store, jwtMid fiber.Handler, web *users.Web, app *fiber.App) {
	router := app.Group("/users")

	router.Post("/registration/admin", web.RegistrationAdminHandler)
	router.Post("/registration/investor", web.RegistrationInvestorHandler)
	router.Post("/registration/company", web.RegistrationCompanyHandler)

	router.Post("/login", web.AuthorizationHandler)
	router.Use(jwtMid, getSessionMiddleware(store)).Get("/profile", web.UserProfileHandler)
}
