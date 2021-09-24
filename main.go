package main

import (
	"github.com/rs/zerolog"
	"sellboot/api"
	"sellboot/migrate"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	migrate.DoMigration()
	api.Start()
}
