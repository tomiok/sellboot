package companies

import (
	"github.com/rs/zerolog/log"
	"sellboot/storage"
)

func DoMigration() {
	db := datastorage.Get()

	err := db.AutoMigrate(
		&Company{},
		&Client{},
		&Investor{},
		&Competitor{},
		&MeasurableAssets{},
		&MidMeasurableAssets{},
		&SoftAssets{},
		&Valuation{},
	)

	if err != nil {
		log.Err(err)
	}
	log.Info().Str("migration", "finish OK")
}
