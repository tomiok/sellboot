package companies

import (
	"github.com/rs/zerolog/log"
	datastorage "sellboot/storage"
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
		log.Error().Msg(err.Error())
	}
	log.Info().Msg("migration finish OK")
}

func Entities() []interface{} {
	return []interface{}{
		&Company{},
		&Client{},
		&Investor{},
		&Competitor{},
		&MeasurableAssets{},
		&MidMeasurableAssets{},
		&SoftAssets{},
		&Valuation{}}
}
