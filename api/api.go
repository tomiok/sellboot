package api

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html"
	"net/http"
	"sellboot/companies"
	"sellboot/configs"
	datastorage "sellboot/storage"
	"sellboot/users"
)

type Server struct {
	*fiber.App
}

func Setup(fs embed.FS) *Server {
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
	return s
}

func newServer(cfg fiber.Config) *Server {
	return &Server{
		App: fiber.New(cfg),
	}
}
