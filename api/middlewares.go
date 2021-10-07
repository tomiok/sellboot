package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"net/http"
	"sellboot/users"
)

func getSessionMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)

		if err != nil {
			return fiber.NewError(http.StatusUnauthorized, "cannot get session")
		}

		log.Info().Msgf("session ID %s", sess.ID())
		log.Info().Msgf("remote IP %s", sess.Get(users.RemoteIP))
		return nil
	}
}
