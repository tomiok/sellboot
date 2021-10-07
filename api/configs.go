package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func appConfig(engine fiber.Views) fiber.Config {
	return fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// retrieve the custom status code if it's a fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Error().Msg(err.Error())
			err = ctx.Status(code).SendFile(fmt.Sprintf("./views/%d.html", code))
			if err != nil {
				log.Error().Msg(err.Error())
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// return from handler
			return nil
		},
	}
}
