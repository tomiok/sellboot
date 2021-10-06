package api

import (
	_ "embed"
	"github.com/gofiber/fiber/v2"
)

func setUpHomeRoutes(app *fiber.App) {
	app.Get("/", homeHandler)
}

func homeHandler(c *fiber.Ctx) error {
	return c.Render("views/templates/home", fiber.Map{"name": "SellBoot"}, "views/layouts/base")
}
