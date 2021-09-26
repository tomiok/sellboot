package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Web struct {
	Svc   UserService
	Store *session.Store
}

type UserRegistration struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"-"`
}

func (w *Web) RegistrationAdminHandler(c *fiber.Ctx) error {
	return addUser(c, w, AdminRole)
}

func (w *Web) RegistrationInvestorHandler(c *fiber.Ctx) error {
	return addUser(c, w, InvestorRole)
}

func (w *Web) RegistrationCompanyHandler(c *fiber.Ctx) error {
	return addUser(c, w, CompanyRole)
}

func (w *Web) AuthorizationHandler(c *fiber.Ctx) error {
	var req LoginRequest

	err := c.BodyParser(&req)

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot parse body")
	}

	login, err := w.Svc.Login(req.Username, req.Password)

	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	//session
	ses, err := w.Store.Get(c)

	if err != nil {
		log.Error().Msgf("cannot create session %s", err.Error())
	}

	ses.Set("token", login.Token)

	return nil
}

func (w *Web) UserProfileHandler(c *fiber.Ctx) error {

	return nil
}

func addUser(c *fiber.Ctx, w *Web, role Role) error {
	var reg UserRegistration
	err := c.BodyParser(&reg)

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot parse body")
	}

	u, err := w.Svc.CreateUser(UserDTO{
		Name:     reg.Name,
		Username: reg.Username,
		Password: reg.Password,
		Role:     role,
	})

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot create user")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": u.ID,
	})
}
