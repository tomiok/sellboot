package main

import (
	"embed"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sellboot/api"
	"sellboot/configs"
	"sellboot/migrate"
)

//go:embed views/*
var viewsFS embed.FS

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	migrate.DoMigration()
	app := api.Setup(viewsFS)
	log.Fatal().Err(app.Listen(":" + configs.Get().Port))
}
