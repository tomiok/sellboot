package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"net/http"
	"sellboot/users"
)

func roleMiddleware(store *session.Store, roles ...users.Role) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)

		if err != nil {
			return fiber.NewError(http.StatusUnauthorized, "cannot get session")
		}

		log.Info().Msgf("session ID %s", sess.ID())
		log.Info().Msgf("remote IP %s", sess.Get(users.RemoteIP))

		b := sess.Get(users.UserData)

		dto := b.(*users.LoginDTO)

		// no roles mean everyone is accepted
		if len(roles) == 0 {
			return nil
		}

		var present = false
		for _, role := range roles {
			if role == dto.Role {
				present = true
				break
			}
		}

		if present {
			return nil
		}

		return fiber.NewError(http.StatusUnauthorized, "you are not allowed to this page")
	}
}
