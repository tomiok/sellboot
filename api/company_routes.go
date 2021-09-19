package api

import (
	"github.com/gofiber/fiber/v2"
	"sellboot/companies"
)

func setUpCompanyRoutes(svc *companies.Web, app *fiber.App) {
	r := app.Group("/companies")

	//views
	r.Get("/", svc.CompanyFormHandler)

	//actions
	r.Get("/valuations/:id", svc.GetValuationHandler)
	r.Get("/rankings", svc.GetCompaniesRankingHandler)

	r.Post("/", svc.SaveCompanyHandler)
}
