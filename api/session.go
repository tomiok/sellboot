package api

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"sellboot/configs"
	"time"
)

func createSession() *session.Store {
	var secure = false

	if configs.Get().Env == "prod" {
		secure = true
	}

	store := session.New(session.Config{
		Expiration:     96 * time.Hour,
		KeyLookup:      "header:session_id",
		CookieSecure:   secure,
		CookieHTTPOnly: false,
		CookieSameSite: "lax",
	})

	return store
}
