package main

import (
	"github.com/rs/zerolog"
	"sellboot/api"
	"sellboot/companies"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	companies.DoMigration()
	api.Start()
}
