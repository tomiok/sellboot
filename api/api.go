package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html"
	"github.com/rs/zerolog/log"
	"sellboot/companies"
	"sellboot/configs"
	datastorage "sellboot/storage"
	"sellboot/users"
	"time"
)

type Server struct {
	*fiber.App
}

func Start() {
	c := configs.Get()
	engine := html.New("./views", ".html")
	_db := datastorage.Get()
	cfg := fiber.Config{
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

	s := newServer(cfg)
	s.Static("/", "./views")
	store := session.New(session.Config{
		Expiration: 96 * time.Hour,
		KeyLookup:  "header:session_id",
	})

	key := configs.Get().JWTSecret
	jwtMid := jwtware.New(jwtware.Config{
		SigningKey: []byte(key),
		ContextKey: configs.Get().JWTContextKey,
	})

	setUpCompanyRoutes(
		store,
		jwtMid,
		&companies.Web{
			Svc: companies.NewGateway(companies.NewStorage(_db)),
		},
		s.App)

	setUpUserRoutes(
		store,
		jwtMid,
		&users.Web{
			Svc:   users.UserService{Gateway: users.NewStorage(_db)},
			Store: store,
		},
		s.App)

	setUpHomeRoutes(s.App)

	port := c.Port

	log.Fatal().Err(s.Listen(":" + port))
}

func newServer(cfg fiber.Config) *Server {
	return &Server{
		App: fiber.New(cfg),
	}
}
