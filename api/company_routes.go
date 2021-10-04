package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"sellboot/companies"
)

func setUpCompanyRoutes(store *session.Store, jwtMid fiber.Handler, svc *companies.Web, app *fiber.App) {
	r := app.Group("/companies")

	//views
	r.Get("/", svc.CompanyFormHandler)

	//actions
	r.Use(jwtMid, getSessionMiddleware(store)).Get("/valuations/:id", svc.GetValuationHandler)
	r.Use(jwtMid).Get("/rankings", svc.GetCompaniesRankingHandler)

	r.Use(jwtMid).Post("/", svc.SaveCompanyHandler)
}
