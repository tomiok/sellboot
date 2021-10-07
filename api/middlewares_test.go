package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/valyala/fasthttp"
	"sellboot/users"
	"testing"
)

func TestRoleMiddleware(t *testing.T) {
	app := fiber.New()

	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	store := session.New(session.Config{})
	store.RegisterType(&users.LoginDTO{})
	sess, err := store.Get(ctx)

	if err != nil {
		t.Fatal()
	}

	sess.Set(users.RemoteIP, "127.0.0.1")
	sess.Set(users.UserData, &users.LoginDTO{
		UserID:    1,
		Status:    "OK",
		Token:     "123asd",
		SessionID: "00001",
		Role:      users.CompanyRole,
	})

	fmt.Println(sess.Save())
	fmt.Println(sess.ID())
	ctx.Request().Header.Set("session_id", sess.ID())
	fn := roleMiddleware(store, users.AdminRole)

	fmt.Println(fn(ctx))
}
