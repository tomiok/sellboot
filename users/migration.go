package users

import (
	"github.com/rs/zerolog/log"
	datastorage "sellboot/storage"
)

func DoMigration() {
	db := datastorage.Get()

	err := db.AutoMigrate(
		&User{},
	)

	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Info().Msg("migration finish OK")
}
