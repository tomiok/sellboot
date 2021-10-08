package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/valyala/fasthttp"
	"sellboot/users"
	"testing"
)

var companyUser = &users.LoginDTO{
	UserID:    1,
	Status:    "OK",
	Token:     "123asd",
	SessionID: "00001",
	Role:      users.CompanyRole,
}

var adminUser = &users.LoginDTO{
	UserID:    2,
	Status:    "OK",
	Token:     "456QWE",
	SessionID: "00002",
	Role:      users.AdminRole,
}

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
	// add an admin role
	sess.Set(users.UserData, adminUser)

	err = sess.Save()
	if err != nil {
		t.Error(err.Error())
	}
	ctx.Request().Header.Set("session_id", sess.ID())

	// add a company role
	fn := roleMiddleware(store, users.CompanyRole)

	// error should occur, if nil, fails here
	if err := fn(ctx); err == nil {
		t.Error("err should not be nil, different roles here")
		t.FailNow()
	}

	// admin role set up
	fn = roleMiddleware(store, users.AdminRole)

	if err := fn(ctx); err != nil {
		t.Error("should be authorized")
		t.FailNow()
	}
}
