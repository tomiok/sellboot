package api

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html"
	"github.com/rs/zerolog/log"
	"net/http"
	"sellboot/companies"
	"sellboot/configs"
	datastorage "sellboot/storage"
	"sellboot/users"
)

type Server struct {
	*fiber.App
}

func Start(fs embed.FS) {
	c := configs.Get()
	engine := html.NewFileSystem(http.FS(fs), ".html")
	_db := datastorage.Get()
	cfg := appConfig(engine)

	s := newServer(cfg)
	s.Static("/", "./views")
	store := createSession()

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
