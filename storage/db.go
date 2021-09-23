package datastorage

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sellboot/configs"
)

var DB *gorm.DB

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

func get(env string) *gorm.DB {
	if env == "local" {
		db, err := gorm.Open(sqlite.Open(configs.Get().DBURL), nil)

		if err != nil {
			log.Fatal().Err(err)
		}
		return db
	}

	if env == "test" {
		db, err := gorm.Open(sqlite.Open(configs.Get().DBURL), nil)

		if err != nil {
			log.Fatal().Err(err)
		}
		return db
	}

	panic("please set a correct env")
}
