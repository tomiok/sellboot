package users

import (
	"os"
	datastorage "sellboot/storage"
	"testing"
)

const (
	pwd = "some_pass_1989$$$"
	jwtSecret = "____________This_"
)

var stg = Storage{
	DB: datastorage.GetTestDB(Entities()...),
}

func Test_Create(t *testing.T) {
	u, err := stg.Create(UserDTO{
		Name:     "tomas",
		Username: "tomasito",
		Password: pwd,
		Role:     AdminRole,
	})

	if err != nil {
		t.Fatal("cannot save user " + err.Error())
	}

	if u.ID == 0 {
		t.Fatal("ID must not be 0")
	}

	if u.Password == pwd {
		t.Fatal("password must be encrypted")
	}
}

func Test_Authenticate(t *testing.T) {
	u, err := stg.Create(UserDTO{
		Name:     "tomas",
		Username: "tomasito69",
		Password: pwd,
		Role:     AdminRole,
	})

	if err != nil {
		t.FailNow()
	}
	_ = os.Setenv("JWT_SECRET", jwtSecret)
	dto, err := stg.Authenticate("tomasito69", pwd)

	if err != nil {
		t.Fatal(err.Error())
	}

	if dto.Status != "OK" {
		t.Fatal("status should be OK")
	}

	if u.ID != dto.UserID {
		t.Fatalf("wanted: %d, got: %d", u.ID, dto.UserID)
	}
}
