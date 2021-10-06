package main

import (
	"embed"
	"github.com/rs/zerolog"
	"sellboot/api"
	"sellboot/migrate"
)

//go:embed views/*
var viewsFS embed.FS

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	migrate.DoMigration()
	api.Start(viewsFS)
}
