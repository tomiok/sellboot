package migrate

import (
	"sellboot/companies"
	"sellboot/users"
)

func DoMigration() {
	companies.DoMigration()
	users.DoMigration()
}

func Entities() []interface{} {
	var res []interface{}

	res = append(res, companies.Entities()...)
	res = append(res, users.Entities()...)

	return res
}
