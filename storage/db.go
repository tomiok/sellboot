package datastorage

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sellboot/configs"
)

var DB *gorm.DB
var DBTest *gorm.DB

func Get() *gorm.DB {
	if DB == nil {
		env := configs.Get().Env
		if env == "" {
			env = "local"
		}
		DB = get(env)
	}
	return DB
}

func GetTestDB(entities ...interface{}) *gorm.DB {
	if DBTest == nil {
		db, err := gorm.Open(sqlite.Open(configs.Get().DBTest), nil)
		if err != nil {
			panic("cannot init TestDB")
		}
		DBTest = db
	}
	err := DBTest.Migrator().DropTable(entities...)

	if err != nil {
		log.Error().Msgf("cannot drop tables %s", err.Error())
	}
	err = DBTest.Migrator().AutoMigrate(entities...)

	if err != nil {
		log.Error().Msgf("cannot drop tables %s", err.Error())
	}

	return DBTest
}

func get(env string) *gorm.DB {
	if env == "local" {
		db, err := gorm.Open(sqlite.Open(configs.Get().DBURL), nil)

		if err != nil {
			log.Fatal().Err(err)
		}
		return db
	}

	panic("please set a correct env")
}
