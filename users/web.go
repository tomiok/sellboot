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

func (w *Web) getSession(c *fiber.Ctx) *session.Session {
	sess, err := w.Store.Get(c)

	if err != nil {
		log.Error().Msgf("cannot get session, %s", err.Error())
		return nil
	}

	return sess
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
	sess := w.getSession(c)

	if sess == nil {
		return fiber.NewError(http.StatusBadRequest, "cannot get session")
	}

	sess.Set(Token, login.Token)
	sess.Set(RemoteIP, c.IP())
	sess.SetExpiry(w.Store.Expiration)
	sess.Set(UserData, login.marshall())

	login.SessionID = sess.ID()
	err = sess.Save()

	if err != nil {
		return err
	}

	return c.JSON(login)
}

func (w *Web) LogoutHandler(c *fiber.Ctx) error {
	sess := w.getSession(c)
	if sess == nil {
		return fiber.NewError(http.StatusBadRequest, "cannot get session")
	}

	return sess.Destroy()
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
