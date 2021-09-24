package users

import (
	datastorage "sellboot/storage"
	"testing"
)

const pwd = "some_pass_1989$$$"

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
