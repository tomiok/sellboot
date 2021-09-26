package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	var reg UserRegistration
	err := c.BodyParser(&reg)

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot parse body")
	}

	u, err := w.Svc.CreateUser(UserDTO{
		Name:     reg.Name,
		Username: reg.Username,
		Password: reg.Password,
		Role:     AdminRole,
	})

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "cannot create user")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": u.ID,
	})
}
