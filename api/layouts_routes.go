package api

import "github.com/gofiber/fiber/v2"

func setUpHomeRoutes(app *fiber.App) {
	app.Get("/", homeHandler)
}

func homeHandler(c *fiber.Ctx) error {
	return c.Render("templates/home", fiber.Map{"name": "SellBoot"}, "layouts/base")
}
